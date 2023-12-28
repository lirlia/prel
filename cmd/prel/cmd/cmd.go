package cmd

import (
	"context"
	"log/slog"
	"os"
	"prel/internal/server"

	"github.com/cockroachdb/errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "prel",
	Short: "prel is a google iam role management system",
	Long: `prel is an application that temporarily assigns Google Cloud IAM Roles and includes an approval process
Complete documentation is available at https://github.com/lirlia/prel`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := checkRequiredFlags(); err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		ctx := context.Background()
		server.Run(ctx)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
func init() {
	cobra.OnInitialize()
	rootCmd.Flags().String("project_id", "", "google cloud project id")
	rootCmd.Flags().String("address", "0.0.0.0", "listen address")
	rootCmd.Flags().Int("port", 8181, "listen port")
	rootCmd.Flags().String("url", "http://localhost:8181", "application url (for redirect)")

	// db
	rootCmd.Flags().String("db_host", "localhost", "postgresql database host")
	rootCmd.Flags().Int("db_port", 5432, "postgresql database port")
	rootCmd.Flags().String("db_user", "postgres", "database user")
	rootCmd.Flags().String("db_password", "", "database password")
	rootCmd.Flags().String("db_name", "prel", "database name")
	rootCmd.Flags().Bool("db_ssl_mode", false, "database ssl mode")
	rootCmd.Flags().String("db_instance_connection", "", "cloud sql connector instance connection name")
	rootCmd.Flags().String("db_type", "fixed", "database type (fixed or cloud_sql_connector)")

	// notification
	rootCmd.Flags().String("notification_type", "slack", "notification type (slack)")
	rootCmd.Flags().String("notification_url", "", "notification url")

	// oauth2
	rootCmd.Flags().String("client_id", "", "google oauth2 client id")
	rootCmd.Flags().String("client_secret", "", "google oauth2 client secret")

	// session
	rootCmd.Flags().Int("session_expire_seconds", 43200, "session expire seconds")

	// debug
	rootCmd.Flags().Bool("is_debWug", false, "debug mode")
	rootCmd.Flags().Bool("is_e2e_mode", false, "e2e test mode")

	viper.AutomaticEnv()
	viper.BindPFlags(rootCmd.Flags())
}

func checkRequiredFlags() error {
	requiredFlags := []string{"project_id", "client_id", "client_secret", "db_password"}
	for _, flag := range requiredFlags {
		if !viper.IsSet(flag) {
			return errors.New("required flag is not set: " + flag)
		}
	}
	return nil
}
