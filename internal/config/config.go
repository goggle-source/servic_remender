package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env      string `mapstructure:"env"`
	GRPC     GRPC   `mapstructure:"grpc"`
	Database DB     `mapstructure:"database"`
}

type DB struct {
	Name     string `mapstructure:"user"`
	Port     int    `mapstructure:"port"`
	DbName   string `mapstructure:"db_name"`
	Password string `mapstructure:"pass"`
}

type GRPC struct {
	Port    int           `mapstructure:"port"`
	Timeout time.Duration `mapstructure:"timeout"`
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	var cfg Config

	if err := viper.ReadInConfig(); err != nil {
		panic("configuration initialization error" + err.Error())
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic("configuration initialization error" + err.Error())
	}

	return &cfg
}
