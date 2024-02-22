package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
	"github.com/sunitha/wheels-away-iam/pkg/gorm"
	"github.com/sunitha/wheels-away-iam/pkg/zitadel"
)

type Config struct {
	Env           string         `mapstructure:"ENVIRONMENT"`
	APIPort       int            `mapstructure:"API_PORT"`
	GRPCPort      int            `mapstructure:"GRPC_PORT"`
	Database      gorm.Config    `mapstructure:",squash"`
	LogLevel      string         `mapstructure:"LOG_LEVEL"`
	LogFormat     string         `mapstructure:"LOG_FORMAT"`
	ZitadelConfig zitadel.Config `mapstructure:",squash"`
}

func Init(configName string) *Config {
	cf := strings.TrimRight(configName, ".")
	viper.SetConfigName(cf)
	viper.AddConfigPath("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
	}

	return NewConfig()
}

func NewConfig() *Config {
	config := Config{}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
	}
	return &config
}
