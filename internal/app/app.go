package app

import (
	"os/signal"
	"syscall"
	"context"
	"os"

	"github.com/nikitaSstepanov/url-shortener/internal/controller/http/v1"
	"github.com/nikitaSstepanov/url-shortener/internal/usecase/storage"
	"github.com/nikitaSstepanov/url-shortener/pkg/client/postgresql"
	"github.com/nikitaSstepanov/url-shortener/internal/usecase"
	"github.com/nikitaSstepanov/url-shortener/pkg/client/redis"
	"github.com/nikitaSstepanov/url-shortener/pkg/logging"
	"github.com/nikitaSstepanov/url-shortener/pkg/server"
	"github.com/nikitaSstepanov/url-shortener/migrations"
)

type App struct {
	controller *controller.Controller
	usecase    *usecase.UseCase
	storage    *storage.Storage
	logger     *logging.Logger
	server     *server.Server
}

func New() *App {
	cfg, err := getAppConfig()
	if err != nil {
		panic("Can`t get app config. Error: " + err.Error())
	}

	logger := logging.NewLogger(&cfg.Logger)

	ctx := context.TODO()

	pgPool, err := postgresql.ConnectToDB(ctx, &cfg.Postgres)
	if err != nil {
		logger.Error("Can`t connect to postgres. Error: " + err.Error())
	} else {
		logger.Info("Postgres is connected")
	}
 
	if err := migrations.Migrate(pgPool); err != nil {
		logger.Error("Can`t migrate postgres scheme. Error: " + err.Error())
	} else {
		logger.Info("Postgres scheme is migrated")
	}

	redisClient, err := redis.ConnectToRedis(ctx, &cfg.Redis)
	if err != nil {
		logger.Error("Can`t connect to redis. Error: " + err.Error())
	} else {
		logger.Info("Redis is connected")
	}

	app := &App{}

	app.logger = logger

	app.storage = storage.New(pgPool, redisClient)

	app.usecase = usecase.New(app.storage)

	app.controller = controller.New(app.usecase)

	handler := app.controller.InitRoutes()

	app.server = server.New(handler, &cfg.Server)

	return app
}

func (a *App) Run() {
	if err := a.server.Start(a.logger); err != nil {
		a.logger.Error("Can`t run application. Error: " + err.Error())
	}

	a.logger.Info("Application is running")

	if err := ShutdownApp(a); err != nil {
		a.logger.Error("Application shutdown error. Error: " + err.Error())
	}

	a.logger.Info("Application Shutting down")
}

func ShutdownApp(a *App) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := a.server.Shutdown(context.Background()); err != nil {
		return err
	}

	return nil
}