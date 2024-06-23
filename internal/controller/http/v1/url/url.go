package url

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikitaSstepanov/url-shortener/internal/controller/http/v1/dto"
	"github.com/nikitaSstepanov/url-shortener/internal/entity"
)

type UrlsUseCase interface {
	SetUrl(ctx context.Context, url *entity.Url) (*entity.Url, *entity.Message)
	GetUrl(ctx context.Context, alias string) (*entity.Url, *entity.Message)
}

type Url struct {
	usecase UrlsUseCase
}

func New(uc UrlsUseCase) *Url {
	return &Url{
		usecase: uc,
	}
}

func (u *Url) SetUrl(ctx echo.Context) error {
	newUrl := dto.SetUrlDto{}

	if err := ctx.Bind(&newUrl); err != nil {
		return ctx.JSON(http.StatusBadRequest, "incorrect body")
	}

	url, msg := u.usecase.SetUrl(context.Background(), newUrl.ToEntity())
	if msg != nil {
		return u.handleMsg(ctx, msg)
	}

	return ctx.JSON(http.StatusOK, dto.UrlToDto(url))
}

func (u *Url) Redirect(ctx echo.Context) error {
	alias := ctx.Param("alias")

	url, msg := u.usecase.GetUrl(context.Background(), alias)
	if msg != nil {
		return u.handleMsg(ctx, msg)
	}

	return ctx.Redirect(http.StatusFound, url.Url)
}

func (u *Url) handleMsg(ctx echo.Context,msg *entity.Message) error {
	e := dto.Error(msg.Msg)

	switch msg.Status {

	case entity.Internal:
		return ctx.JSON(http.StatusInternalServerError, e)

	case entity.Conflict:
		return ctx.JSON(http.StatusConflict, e)

	case entity.NotFound:
		return ctx.JSON(http.StatusNotFound, e)
	
	case entity.BadInput:
		return ctx.JSON(http.StatusBadRequest, e)

	}

	return nil
}