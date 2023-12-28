package config

const AppName = "prel"
const RedirectPath = "/auth/google/callback"

type Config struct {
	// ProjectID is the GCP project ID
	ProjectID string `mapstructure:"project_id"`
	Address   string `mapstructure:"address"`
	Port      string `mapstructure:"port"`

	// URL is the base URL of the application
	URL string `mapstructure:"url"`

	// for database(postgresql)
	DBUsername string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBName     string `mapstructure:"db_name"`
	DBHost     string `mapstructure:"db_host"`
	DBPort     int    `mapstructure:"db_port"`
	DBSslMode  bool   `mapstructure:"db_ssl_mode"`

	// DBInstanceConnection is the connection name for Cloud SQL Connector
	DBInstanceConnection string `mapstructure:"db_instance_connection"`

	// DBType can be "fixed" or "cloud_sql_connector"
	// "fixed" means that the database is fixed to a specific host and port
	// "cloud_sql_connector" means that the database is connected via Cloud SQL Connector
	DBType string `mapstructure:"db_type"`

	// NotificationType can be "slack"
	NotificationType string `mapstructure:"notification_type"`

	// NotificationUrl is the URL to send notification to
	// For Slack, it is the webhook URL
	NotificationUrl string `mapstructure:"notification_url"`

	// ClientID and ClientSecret are the Google OAuth2 client ID and secret
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`

	// SessionExpireSeconds is the number of seconds until the session expires
	// Default is 12 hours
	SessionExpireSeconds int `mapstructure:"session_expire_seconds"`

	// IsDebug is the flag to enable debug mode
	// If true, the application will output request payload, and debug logs
	IsDebug bool `mapstructure:"is_debug"`

	//
	// CAUTION: DO NOT USE THIS FLAG IN PRODUCTION
	//
	// IsE2EMode is the flag to enable e2e test mode
	// If true, the application will use a suitable setting
	IsE2EMode bool `mapstructure:"is_e2e_mode"`
}

var isDebug bool

func SetDebug(debug bool) {
	isDebug = debug
}

func IsDebug() bool {
	return isDebug
}
