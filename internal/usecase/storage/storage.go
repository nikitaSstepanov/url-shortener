package storage

import (
	"github.com/nikitaSstepanov/url-shortener/internal/usecase/storage/url"
	"github.com/nikitaSstepanov/url-shortener/pkg/client/postgresql"
	goredis "github.com/redis/go-redis/v9"
)

type Storage struct {
	Url *url.UrlStorage
}

func New(pg postgresql.Client, redis *goredis.Client) *Storage {
	return &Storage{
		Url: url.New(pg, redis),
	}
}