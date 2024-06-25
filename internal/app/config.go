package app

import (
	"github.com/nikitaSstepanov/url-shortener/pkg/client/postgresql"
	"github.com/nikitaSstepanov/url-shortener/pkg/client/redis"
	"github.com/nikitaSstepanov/url-shortener/pkg/logging"
	"github.com/nikitaSstepanov/url-shortener/pkg/server"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Server   server.Config     `yaml:"server"`
	Logger   logging.Config    `yaml:"logger"`
	Postgres postgresql.Config `yaml:"postgres"`
	Redis    redis.Config      `yaml:"redis"`
}

func getAppConfig() (*AppConfig, error) {
	var cfg AppConfig

	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	err := cleanenv.ReadConfig("config/config.yaml", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}