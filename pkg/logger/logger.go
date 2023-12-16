package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
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

func NewLogger(config *config.Config) *zerolog.Logger {
	level, ok := errorLevelMap[config.LogLevel]
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: true}
	output.FormatCaller = func(i interface{}) string {
		return fmt.Sprintf("caller:%s | ", i)
	}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("message: %s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("| %s: ", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s |", i)
	}

	if !ok {
		log.Panicf("%s loglevel is invalid", config.LogLevel)
	}
	logger := zerolog.New(output).Level(level).With().Logger()
	logger = logger.With().Caller().Logger()
	logger = logger.With().Timestamp().Logger()
	return &logger
}
