package dto

import "github.com/nikitaSstepanov/url-shortener/internal/entity"

type SetUrlDto struct {
	Url   string `json:"url"`
	Alias string `json:"alias"`
}

type SendAliasDto struct {
	Alias string `json:"alias"`
}

func (u *SetUrlDto) ToEntity() *entity.Url {
	return &entity.Url{
		Url:   u.Url,
		Alias: u.Alias,
	}
}

func UrlToDto(u *entity.Url) *SendAliasDto {
	return &SendAliasDto{
		Alias: u.Alias,
	}
}