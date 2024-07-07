package url

import (
	"context"
	
	"github.com/nikitaSstepanov/url-shortener/pkg/utils/generator"
	"github.com/nikitaSstepanov/url-shortener/internal/entity"
)

type UrlStorage interface {
	GetUrl(ctx context.Context, alias string) (*entity.Url, error)
	SetUrl(ctx context.Context, url *entity.Url) error
}

type UrlUseCase struct {
	urls UrlStorage
}

const (
	aliasLength int = 6
)

func New(store UrlStorage) *UrlUseCase {
	return &UrlUseCase{
		urls: store,
	}
}

func (u *UrlUseCase) SetUrl(ctx context.Context, url *entity.Url) (*entity.Url, *entity.Message) {
	if url.Url == "" {
		return nil, entity.GetMsg("Url mustn`t be empty", entity.BadInput)
	}
	
	if url.Alias == "" {
		for {
			url.Alias = generator.GetRandomString(aliasLength)

			candidate, err := u.urls.GetUrl(ctx, url.Alias)
			if err != nil {
				return nil, entity.GetMsg("Something going wrong...", entity.Internal)
			}

			if candidate.Id == 0 {
				break
			}
		}
	} else {
		candidate, err := u.urls.GetUrl(ctx, url.Alias)
		if err != nil {
			return nil, entity.GetMsg("Something going wrong...", entity.Internal)
		}

		if candidate.Id != 0 {
			return nil, entity.GetMsg("This alias was taken :(", entity.Conflict)
		}
	}

	if err := u.urls.SetUrl(ctx, url); err != nil {
		return nil, entity.GetMsg("Something going wrong...", entity.Internal)
	}

	return url, nil
}

func (u *UrlUseCase) GetUrl(ctx context.Context, alias string) (*entity.Url, *entity.Message) {
	url, err := u.urls.GetUrl(context.Background(), alias)
	if err != nil {
		return nil, entity.GetMsg("Something going wrong..", entity.Internal)
	}

	if url.Id == 0 {
		return nil, entity.GetMsg("This url wasn`t found", entity.NotFound)
	}

	return url, nil
}