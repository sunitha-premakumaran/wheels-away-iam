package gorm

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	logger   *zerolog.Logger
	logLevel LogLevel
}

const (
	SourceKey   = "caller"
	ErrorKey    = "error"
	QueryKey    = "query"
	DurationKey = "duration"
	RowsKey     = "rows_affected"
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
	elapsed := time.Since(begin)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		sql, rows := fc()

		g.logger.Err(err).Ctx(ctx).Str(QueryKey, sql).Dur(DurationKey, elapsed).Int64(RowsKey, rows)
	} else if g.logLevel == Info {
		sql, rows := fc()

		g.logger.Info().Ctx(ctx).Str(QueryKey, sql).Dur(DurationKey, elapsed).Int64(RowsKey, rows)
	}
}

func NewGormLogger(logger *zerolog.Logger, config *Config) *GormLogger {
	return &GormLogger{
		logger:   logger,
		logLevel: config.LogLevel,
	}
}
