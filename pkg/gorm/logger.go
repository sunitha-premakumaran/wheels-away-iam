package gorm

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type GormLogger struct {
	ignoreTrace               bool
	ignoreRecordNotFoundError bool
	traceAll                  bool
	slowThreshold             time.Duration

	sourceField string
	errorField  string
	logger      *zerolog.Logger
}

type LogType string

const (
	ErrorLogType     LogType = "sql_error"
	SlowQueryLogType LogType = "slow_query"
	DefaultLogType   LogType = "default"

	SourceField    = "file"
	ErrorField     = "error"
	QueryField     = "query"
	DurationField  = "duration"
	SlowQueryField = "slow_query"
	RowsField      = "rows"
)

func (g *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return g
}

func (g *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	g.logger.Info().Fields(data).Msg(msg)
}

func (g *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	g.logger.Warn().Fields(data).Msg(msg)
}

func (g *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	g.logger.Error().Fields(data).Msg(msg)
}

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.ignoreTrace {
		return // Silent
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && (!errors.Is(err, gorm.ErrRecordNotFound) || !g.ignoreRecordNotFoundError):
		sql, rows := fc()

		g.logger.Err(err).Ctx(ctx).Str(QueryField, sql).Dur(DurationField, elapsed).Int64(RowsField, rows).Str(g.sourceField, utils.FileWithLineNum())

	case g.slowThreshold != 0 && elapsed > g.slowThreshold:
		sql, rows := fc()

		g.logger.Info().Ctx(ctx).Bool(SlowQueryField, true).Str(QueryField, sql).
			Dur(DurationField, elapsed).Int64(RowsField, rows).Str(g.sourceField, utils.FileWithLineNum()).
			Msgf("slow sql query [%s >= %v]", elapsed, g.slowThreshold)

	case g.traceAll:
		sql, rows := fc()

		g.logger.Info().Ctx(ctx).Str(QueryField, sql).Dur(DurationField, elapsed).Int64(RowsField, rows).Str(g.sourceField, utils.FileWithLineNum()).
			Msgf("SQL query executed [%s]", elapsed)
	}
}

func NewGormLogger(logger *zerolog.Logger, config *Config) *GormLogger {
	return &GormLogger{
		logger:                    logger,
		ignoreTrace:               config.LogLevel == "SILENT",
		ignoreRecordNotFoundError: true,
		traceAll:                  true,
		slowThreshold:             9500000,
	}
}
