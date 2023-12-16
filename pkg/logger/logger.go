package logger

import (
	"bytes"
	"log"
	"os"
	"time"

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

func NewLogger(app string, config *config.Config) *zerolog.Logger {
	level, ok := errorLevelMap[config.LogLevel]
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: true}
	output.FormatExtra = func(m map[string]interface{}, b *bytes.Buffer) error {
		// keys := make([]string, 0, len(m))
		// for k := range m {
		// 	keys = append(keys, k)
		// }
		return nil
	}
	zerolog.TimestampFieldName = "timestamp"
	if !ok {
		log.Panicf("%s loglevel is invalid", config.LogLevel)
	}
	logger := zerolog.New(output).Level(level).With().Str("app", app).Logger()
	logger = logger.With().Caller().Logger()
	logger = logger.With().Timestamp().Logger()
	return &logger
}
