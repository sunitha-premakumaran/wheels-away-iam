package gorm

type Config struct {
	Server                 string   `mapstructure:"DATABASE_SERVER"`
	Port                   int      `mapstructure:"DATABASE_PORT"`
	Name                   string   `mapstructure:"DATABASE_NAME"`
	User                   string   `mapstructure:"DATABASE_USER"`
	Password               string   `mapstructure:"DATABASE_PASSWORD"`
	LogLevel               LogLevel `mapstructure:"DATABASE_LOG_LEVEL"`
	SslMode                SslMode  `mapstructure:"DATABASE_SSL_MODE"`
	NewRelicTracingEnabled bool     `mapstructure:"NEW_RELIC_TRACING_ENABLED"`
}

type LogLevel string

const (
	Info  LogLevel = "INFO"
	Error LogLevel = "ERROR"
)

type SslMode string

const (
	SslModeDisable    SslMode = "disable"
	SslModeAllow      SslMode = "allow"
	SslModePrefer     SslMode = "prefer"
	SslModeRequire    SslMode = "require"
	SslModeVerifyCa   SslMode = "verify-ca"
	SslModeVerifyFull SslMode = "verify-full"
)
