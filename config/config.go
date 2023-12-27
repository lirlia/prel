package config

const AppName = "prel"

type Config struct {
	// ProjectID is the GCP project ID
	ProjectID string `envconfig:"PROJECT_ID"`
	Address   string `envconfig:"ADDRESS" default:"0.0.0.0"`
	Port      string `envconfig:"PORT" default:"8181"`

	// URL is the base URL of the application
	URL string `envconfig:"URL" default:"http://localhost:8181"`

	// for database(postgresql)
	DBUsername string `envconfig:"DB_USERNAME" default:"postgres"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"password"`
	DBName     string `envconfig:"DB_NAME" default:"prel"`
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     int    `envconfig:"DB_PORT" default:"5432"`
	DBSslMode  bool   `envconfig:"DB_SSL_MODE" default:"false"`

	// DBInstanceConnection is the connection name for Cloud SQL Connector
	DBInstanceConnection string `envconfig:"DB_INSTANCE_CONNECTION" default:""`

	// DBType can be "fixed" or "cloud_sql_connector"
	// "fixed" means that the database is fixed to a specific host and port
	// "cloud_sql_connector" means that the database is connected via Cloud SQL Connector
	DBType string `envconfig:"DB_TYPE" default:"fixed"`

	// NotificationType can be "slack"
	NotificationType string `envconfig:"NOTIFICATION_TYPE" default:"slack"`

	// NotificationUrl is the URL to send notification to
	// For Slack, it is the webhook URL
	NotificationUrl string `envconfig:"NOTIFICATION_URL"`

	// RedirectPath is the path to redirect after authentication
	RedirectPath string `envconfig:"REDIRECT_URL" default:"/auth/google/callback"`

	// ClientID and ClientSecret are the Google OAuth2 client ID and secret
	ClientID     string `envconfig:"CLIENT_ID"`
	ClientSecret string `envconfig:"CLIENT_SECRET"`

	// SessionExpireSeconds is the number of seconds until the session expires
	// Default is 12 hours
	SessionExpireSeconds int `envconfig:"SESSION_EXPIRE_SECONDS" default:"43200"`

	// IsDebug is the flag to enable debug mode
	// If true, the application will output request payload, and debug logs
	IsDebug bool `envconfig:"IS_DEBUG" default:"false"`

	//
	// CAUTION: DO NOT USE THIS FLAG IN PRODUCTION
	//
	// IsE2EMode is the flag to enable e2e test mode
	// If true, the application will use a suitable setting
	IsE2EMode bool `envconfig:"IS_E2E_MODE" default:"false"`
}

var isDebug bool

func SetDebug(debug bool) {
	isDebug = debug
}

func IsDebug() bool {
	return isDebug
}
