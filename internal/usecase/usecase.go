package usecase

import (
	"github.com/nikitaSstepanov/url-shortener/internal/usecase/storage"
	"github.com/nikitaSstepanov/url-shortener/internal/usecase/url"
)

type UseCase struct {
	Url *url.Url
}

func New(store *storage.Storage) *UseCase {
	return &UseCase{
		Url: url.New(store.Url),
	}
}