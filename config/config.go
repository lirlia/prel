package config

const AppName = "prel"
const RedirectPath = "/auth/google/callback"

type Config struct {
	// ProjectID is the GCP project ID
	ProjectID string `mapstructure:"project-id"`
	Address   string `mapstructure:"address"`
	Port      string `mapstructure:"port"`

	// URL is the base URL of the application
	URL string `mapstructure:"url"`

	// for database(postgresql)
	DBUsername string `mapstructure:"db-user"`
	DBPassword string `mapstructure:"db-password"`
	DBName     string `mapstructure:"db-name"`
	DBHost     string `mapstructure:"db-host"`
	DBPort     int    `mapstructure:"db-port"`
	DBSslMode  bool   `mapstructure:"db-ssl-mode"`

	// DBInstanceConnection is the connection name for Cloud SQL Connector
	DBInstanceConnection string `mapstructure:"db-instance-connection"`

	// DBType can be "fixed" or "cloud-sql-connector"
	// "fixed" means that the database is fixed to a specific host and port
	// "cloud-sql-connector" means that the database is connected via Cloud SQL Connector
	DBType DBtype `mapstructure:"db-type"`

	// NotificationType can be "slack"
	NotificationType NotificationType `mapstructure:"notification-type"`

	// NotificationUrl is the URL to send notification to
	// For Slack, it is the webhook URL
	NotificationUrl string `mapstructure:"notification-url"`

	AuthnType AuthnType `mapstructure:"authentication-type"`
	// ClientID and ClientSecret are the Google OAuth2 client ID and secret
	ClientID     string `mapstructure:"client-id"`
	ClientSecret string `mapstructure:"client-secret"`
	IapAudience  string `mapstructure:"iap-audience"`

	// SessionExpireSeconds is the number of seconds until the session expires
	// Default is 12 hours
	SessionExpireSeconds int `mapstructure:"session-expire-seconds"`

	// RequestExpireSeconds is the number of seconds until the request expires
	// Default is 24 hours
	RequestExpireSeconds int `mapstructure:"request-expire-seconds"`

	// IsDebug is the flag to enable debug mode
	// If true, the application will output request payload, and debug logs
	IsDebug bool `mapstructure:"is-debug"`

	//
	// CAUTION: DO NOT USE THIS FLAG IN PRODUCTION
	//
	// IsE2EMode is the flag to enable e2e test mode
	// If true, the application will use a suitable setting
	IsE2EMode bool `mapstructure:"is-e2e-mode"`
}

var isDebug bool

func SetDebug(debug bool) {
	isDebug = debug
}

func IsDebug() bool {
	return isDebug
}

type DBtype string

const (
	FixedDB           DBtype = "fixed"
	CloudSQLConnector DBtype = "cloud-sql-connector"
)

type NotificationType string

const (
	Slack NotificationType = "slack"
)

type AuthnType string

const (
	Google AuthnType = "google"
	IAP    AuthnType = "iap"
)
