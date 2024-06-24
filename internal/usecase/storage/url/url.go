package url

import (
	"context"
	"time"
	"fmt"

	"github.com/nikitaSstepanov/url-shortener/internal/usecase/storage/model"
	"github.com/nikitaSstepanov/url-shortener/pkg/client/postgresql"
	"github.com/nikitaSstepanov/url-shortener/internal/entity"
	goredis "github.com/redis/go-redis/v9"
	"github.com/jackc/pgx/v5"
)

type Url struct {
	postgres postgresql.Client
	redis    *goredis.Client
}

const (
	redisExpires = 3 * time.Hour
	urlsTable    = "urls"
)

func New(pg postgresql.Client, redisConn *goredis.Client) *Url {
	return &Url{
		postgres: pg,
		redis:    redisConn,
	}
}

func (u *Url) GetUrl(ctx context.Context, alias string) (*entity.Url, error) {
	var url model.Url

	err := u.redis.Get(ctx, getRedisKey(alias)).Scan(&url)
	if err != nil && err != goredis.Nil {
		return nil, err
	}

	if url.Id != 0 {
		return url.ToEntity(), nil
	}

	query := fmt.Sprintf("SELECT id, url FROM %s WHERE alias = '%s'", urlsTable, alias)

	row := u.postgres.QueryRow(ctx, query)

	err = row.Scan(&url.Id, &url.Url)
	if err != nil  {
		if err != pgx.ErrNoRows {
			return nil, err
		} else {
			return &entity.Url{}, nil
		}
	}
	
	err = u.redis.Set(ctx, getRedisKey(alias), url, redisExpires).Err()
	if err != nil {
		return nil, err
	}

	return url.ToEntity(), nil
}

func (u *Url) SetUrl(ctx context.Context, url *entity.Url) error {
	urlModel := model.UrlToModel(url)
	
	query := "INSERT INTO urls (url, alias) VALUES ($1, $2) RETURNING id;"

	row := u.postgres.QueryRow(ctx, query, urlModel.Url, urlModel.Alias)
	
	err := row.Scan(&urlModel.Id)
	if err != nil {
		return err
	}

	err = u.redis.Set(ctx, getRedisKey(urlModel.Alias), urlModel, redisExpires).Err()
	if err != nil {
		return err
	}

	return nil
}

func getRedisKey(alias string) string {
	return fmt.Sprintf("urls:%s", alias)
}