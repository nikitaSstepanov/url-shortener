package controller

import (
	"github.com/labstack/echo/v4"
	url_handler "github.com/nikitaSstepanov/url-shortener/internal/controller/http/v1/url"
	"github.com/nikitaSstepanov/url-shortener/internal/usecase"
)

type Controller struct {
	url *url_handler.Url
}

func New(uc *usecase.UseCase) *Controller {
	return &Controller{
		url: url_handler.New(uc.Url),
	}
}

func (c *Controller) InitRoutes() *echo.Echo {
	router := echo.New()

	router.GET("/:alias", c.url.Redirect)

	api := router.Group("/api/v1") 
	{
		url := api.Group("/url") 
		{
			url.POST("/new", c.url.SetUrl)
		}
	}

	return router
}