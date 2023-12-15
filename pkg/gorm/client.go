package gorm

import (
	"fmt"

	"github.com/rs/zerolog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient struct {
	DB *gorm.DB
}

func NewDBClient(config *Config, logger *zerolog.Logger) *DBClient {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Kolkata",
		config.Server, config.User, config.Password, config.Name, config.Port, config.SslMode)
	gormDB, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: NewGormLogger(logger, config),
	})
	if err != nil {
		logger.Panic().Msgf("Connection to database failed %s", err.Error())
	}
	logger.Info().Msgf("connection to database: %s successful!", config.Name)
	return &DBClient{
		DB: gormDB,
	}
}
