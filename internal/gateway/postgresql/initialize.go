package postgresql

import (
	"context"
	"fmt"
	"net"
	"sync"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	query *Queries
	conn  *pgxpool.Pool
	once  sync.Once
)

type Option func(user, password, dbName string) (*pgxpool.Pool, error)

func WithCloudSQLConnector(instanceConnectionName string) Option {
	return func(user, password, dbName string) (*pgxpool.Pool, error) {
		dsn := fmt.Sprintf("user=%s password=%s database=%s", user, password, dbName)
		config, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			return nil, err
		}

		d, err := cloudsqlconn.NewDialer(context.Background())
		if err != nil {
			return nil, err
		}

		config.ConnConfig.DialFunc = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return d.Dial(ctx, instanceConnectionName)
		}

		pool, err := pgxpool.NewWithConfig(context.Background(), config)
		return pool, err
	}
}

func WithFixedDB(host string, port int, sslMode bool) Option {
	return func(user, password, dbName string) (*pgxpool.Pool, error) {
		var ssl string
		if sslMode {
			ssl = "enable"
		} else {
			ssl = "disable"
		}

		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, ssl)
		pool, err := pgxpool.New(context.Background(), dsn)
		return pool, err
	}
}

func Initialize(user, password, dbName string, option Option) error {
	var err error

	once.Do(func() {
		conn, err = option(user, password, dbName)
		if err != nil {
			return
		}

		err = conn.Ping(context.Background())
		if err != nil {
			return
		}

		query = New(conn)
	})

	return err
}

type queryKey struct{}

func SetQueries(ctx context.Context, queries *Queries) context.Context {
	return context.WithValue(ctx, queryKey{}, queries)
}

func getQueries(ctx context.Context) *Queries {
	queries, ok := ctx.Value(queryKey{}).(*Queries)
	if !ok {
		return nil
	}
	return queries
}

func GetQuery(ctx context.Context) *Queries {
	txq := getQueries(ctx)
	if txq != nil {
		return txq
	}
	return query
}

func GetConn() *pgxpool.Pool {
	return conn
}

// use for test
func ResetOnce() {
	once = sync.Once{}
}
