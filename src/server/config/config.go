package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Mongo struct {
	User     string `env:"mongo_user"`
	Pass     string `env:"mongo_pass"`
	Host     string `env:"mongo_host"`
	Args     string `env:"mongo_args"`
	Database string `env:"mongo_database"`
}

type Application struct {
	ReadTimeout  time.Duration `env:"app_readTimeout"`
	WriteTimeout time.Duration `env:"app_writeTimeout"`
}

type Config struct {
	Environment string
	Mongo       Mongo
	Application Application
}

var Instance *Config = nil

func ReadConfigFromEnv(environment string) (Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return Config{}, fmt.Errorf(`error reading env: %w`, err)
	}

	cfg.Environment = environment

	Instance = &cfg

	return cfg, nil
}
