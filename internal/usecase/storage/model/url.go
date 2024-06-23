package model

import (
	"encoding/json"

	"github.com/nikitaSstepanov/url-shortener/internal/entity"
)

type Url struct {
	Id    uint64 `redis:"id"`
	Url   string `redis:"url"`
	Alias string `redis:"alias"`
}

func (u Url) MarshalBinary() ([]byte, error) {
	return json.Marshal(&u)
}

func (u *Url) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

func (u *Url) ToEntity() *entity.Url {
	return &entity.Url{
		Id:    u.Id,
		Url:   u.Url,
		Alias: u.Alias,
	}
}

func UrlToModel(u *entity.Url) *Url {
	return &Url{
		Id:    u.Id,
		Url:   u.Url,
		Alias: u.Alias,
	}
}