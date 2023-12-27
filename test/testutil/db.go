package testutil

import (
	"context"
	"fmt"
	"prel/config"
	"prel/db"
	compose "prel/docker"
	"prel/internal/gateway/postgresql"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

const (
	postgresUser          = "postgres"
	postgresPassword      = "postgres"
	postgresDefaultDBName = config.AppName
)

var (
	once             sync.Once
	postgresResource *dockertest.Resource
	defaultConn      *pgxpool.Pool
	postgresPort     int
	dockerPool       *dockertest.Pool
)

func SetupTestDB() {
	var err error
	once.Do(func() {
		dockerPool, postgresResource, err = runDockerDB()
		if err != nil {
			panic(err)
		}

		portString := postgresResource.GetPort("5432/tcp")
		postgresPort, err = strconv.Atoi(portString)
		if err != nil {
			panic(err)
		}

		fn := postgresql.WithFixedDB("localhost", postgresPort, false)

		err = dockerPool.Retry(func() error {
			defaultConn, err = fn(postgresUser, postgresPassword, postgresDefaultDBName)
			if err != nil {
				return err
			}
			return defaultConn.Ping(context.Background())
		})

		if err != nil {
			panic(err)
		}
	})

	postgresql.ResetOnce()
	dbName := generateDBName()

	ctx := context.Background()
	err = createDB(ctx, dbName)
	if err != nil {
		panic(err)
	}

	err = postgresql.Initialize(postgresUser, postgresPassword, dbName,
		postgresql.WithFixedDB("localhost", postgresPort, false))

	if err != nil {
		panic(err)
	}

	err = applyDDL(ctx)
	if err != nil {
		panic(err)
	}
}

func runDockerDB() (*dockertest.Pool, *dockertest.Resource, error) {

	pool, err := dockertest.NewPool("")
	pool.MaxWait = 10 * time.Second
	if err != nil {
		return nil, nil, err
	}

	postgresVersion, err := compose.GetDBVersion()
	if err != nil {
		return nil, nil, err
	}
	runOptions := &dockertest.RunOptions{
		Repository: "postgres",
		Tag:        postgresVersion,
		Env: []string{
			fmt.Sprintf("POSTGRES_USER=%s", postgresUser),
			fmt.Sprintf("POSTGRES_PASSWORD=%s", postgresPassword),
			fmt.Sprintf("POSTGRES_DB=%s", postgresDefaultDBName),
			"POSTGRES_INITDB_ARGS: --encoding=UTF-8 --locale=C",
			"listen_addresses='*'",
		},
	}

	r, err := pool.RunWithOptions(runOptions,
		func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)
	return pool, r, err
}

func generateDBName() string {
	uuid := uuid.New().String()
	uuidWithoutHyphens := strings.Replace(uuid, "-", "", -1)

	return "test_" + uuidWithoutHyphens
}

func createDB(ctx context.Context, name string) error {
	_, err := defaultConn.Exec(ctx, "CREATE DATABASE "+name)
	return err
}

func applyDDL(ctx context.Context) error {
	ddl := string(db.GetSchema())
	_, err := postgresql.GetConn().Exec(ctx, ddl)
	return err
}

func StopTestDB() {
	if postgresResource != nil {
		postgresResource.Close()
	}
	if dockerPool != nil {
		_ = dockerPool.Purge(postgresResource)
	}
}
