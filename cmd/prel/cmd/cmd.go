package cmd

import (
	"context"
	"log/slog"
	"os"
	"prel/config"
	"prel/internal/server"
	"strings"

	"github.com/cockroachdb/errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type runner interface {
	Run(ctx context.Context)
}
type realRunner struct{}

func (r *realRunner) Run(ctx context.Context) {
	server.Run(ctx)
}

type mockRunner struct {
	Fn func()
}

func (r *mockRunner) Run(_ context.Context) {
	r.Fn()
}

var commandRunner runner

var RootCmd = &cobra.Command{
	Use:   "prel",
	Short: "prel is a google iam role management system",
	Long: `prel is an application that temporarily assigns Google Cloud IAM Roles and includes an approval process
Complete documentation is available at https://github.com/lirlia/prel`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := Validate(); err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		ctx := context.Background()
		commandRunner.Run(ctx)
	},
}

func Execute() error {
	commandRunner = &realRunner{}
	return RootCmd.Execute()
}

func ExecuteTest(fn func()) error {
	commandRunner = &mockRunner{
		Fn: fn,
	}
	return RootCmd.Execute()
}

func init() {

	// cmdline flag use hyphen, but os environment use underscore
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	cobra.OnInitialize()
	RootCmd.Flags().String("project-id", "", "google cloud project id")
	RootCmd.Flags().String("address", "0.0.0.0", "listen address")
	RootCmd.Flags().Int("port", 8181, "listen port")
	RootCmd.Flags().String("url", "http://localhost:8181", "application url (for redirect)")

	// db
	RootCmd.Flags().String("db-host", "localhost", "postgresql database host")
	RootCmd.Flags().Int("db-port", 5432, "postgresql database port")
	RootCmd.Flags().String("db-user", "postgres", "database user")
	RootCmd.Flags().String("db-password", "", "database password")
	RootCmd.Flags().String("db-name", "prel", "database name")
	RootCmd.Flags().Bool("db-ssl-mode", false, "database ssl mode")
	RootCmd.Flags().String("db-instance-connection", "", "cloud sql connector instance connection name")
	RootCmd.Flags().String("db-type", "fixed", "database type (fixed or cloud-sql-connector)")

	// notification
	RootCmd.Flags().String("notification-type", "slack", "notification type (slack)")
	RootCmd.Flags().String("notification-url", "", "notification url")

	// google authentication
	RootCmd.Flags().String("authentication-type", "google", "authentication type (google or iap)")
	RootCmd.Flags().String("client-id", "", "google oauth2 client id (need to set if authentication-type is google)")
	RootCmd.Flags().String("client-secret", "", "google oauth2 client secret (need to set if authentication-type is google)")
	RootCmd.Flags().String("iap-audience", "", "iap audience (need to set if authentication-type is iap) see: https://cloud.google.com/iap/docs/signed-headers-howto?hl=ja#verifying_the_jwt_payload")

	// application config
	RootCmd.Flags().Int("session-expire-seconds", 43200, "how long the user login session will expire")
	RootCmd.Flags().Int("request-expire-seconds", 86400, "how long the user request will expire")

	// debug
	RootCmd.Flags().Bool("is-debug", false, "debug mode")
	RootCmd.Flags().Bool("is-e2e-mode", false, "e2e test mode")

	err := viper.BindPFlags(RootCmd.Flags())
	if err != nil {
		panic(err)
	}
}

func Validate() error {
	requiredFlags := []string{"project-id", "db-password"}
	for _, flag := range requiredFlags {
		if !viper.IsSet(flag) {
			return errors.New("required flag is not set: " + flag)
		}
	}

	// authentication type
	authenticationType := viper.GetString("authentication-type")
	if authenticationType != string(config.Google) && authenticationType != string(config.IAP) {
		return errors.New("invalid authentication-type: " + authenticationType)
	}

	switch authenticationType {
	case string(config.Google):
		if !viper.IsSet("client-id") {
			return errors.New("required flag is not set: client-id")
		}
		if !viper.IsSet("client-secret") {
			return errors.New("required flag is not set: client-secret")
		}
	case string(config.IAP):
		if !viper.IsSet("iap-audience") {
			return errors.New("required flag is not set: iap-audience")
		}
	}

	// db
	dbType := viper.GetString("db-type")
	if dbType != string(config.FixedDB) && dbType != string(config.CloudSQLConnector) {
		return errors.New("invalid db-type: " + dbType)
	}

	if dbType == string(config.CloudSQLConnector) && !viper.IsSet("db-instance-connection") {
		return errors.New("required flag is not set: db-instance-connection")
	}

	return nil
}
