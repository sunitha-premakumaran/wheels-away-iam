package logger

import (
	"log"
	"os"

	"github.com/rs/zerolog"
	"github.com/sunitha/wheels-away-iam/config"
)

var (
	errorLevelMap = map[string]zerolog.Level{
		"info":  zerolog.InfoLevel,
		"debug": zerolog.DebugLevel,
		"error": zerolog.ErrorLevel,
		"warn":  zerolog.WarnLevel,
		"trace": zerolog.TraceLevel,
		"":      zerolog.Disabled,
	}
)

func Init(app string, config *config.Config) *zerolog.Logger {
	level, ok := errorLevelMap[config.LogLevel]
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "timestamp"
	if !ok {
		log.Panicf("%s loglevel is invalid", config.LogLevel)
	}
	logger := zerolog.New(os.Stdout).Level(level).With().Str("app", app).Logger()
	logger = logger.With().Caller().Logger()
	logger = logger.With().Timestamp().Logger()
	return &logger
}
