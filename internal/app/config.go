package app

import (
	"github.com/nikitaSstepanov/url-shortener/pkg/client/postgresql"
	"github.com/nikitaSstepanov/url-shortener/pkg/client/redis"
	"github.com/nikitaSstepanov/url-shortener/pkg/logging"
	"github.com/nikitaSstepanov/url-shortener/pkg/server"
	"github.com/ilyakaznacheev/cleanenv"
)

func getPostgresConfig() (*postgresql.Config, error) {
	var cfg postgresql.Config

	err := cleanenv.ReadConfig("config/postgres.yaml", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, err
}

func getRedisConfig() (*redis.Config, error) {
	var cfg redis.Config

	err := cleanenv.ReadConfig("config/redis.yaml", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func getServerConfig() (*server.Config, error) {
	var cfg server.Config

	err := cleanenv.ReadConfig("config/server.yaml", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func getLoggerConfig() (*logging.Config, error) {
	var cfg logging.Config

	err := cleanenv.ReadConfig("config/logger.yaml", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}