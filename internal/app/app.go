package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/nikitaSstepanov/url-shortener/internal/controller/http/v1"
	"github.com/nikitaSstepanov/url-shortener/internal/usecase"
	"github.com/nikitaSstepanov/url-shortener/internal/usecase/storage"
	"github.com/nikitaSstepanov/url-shortener/migrations"
	"github.com/nikitaSstepanov/url-shortener/pkg/client/postgresql"
	"github.com/nikitaSstepanov/url-shortener/pkg/client/redis"
	"github.com/nikitaSstepanov/url-shortener/pkg/logging"
	"github.com/nikitaSstepanov/url-shortener/pkg/server"
)

type App struct {
	controller *controller.Controller
	usecase    *usecase.UseCase
	storage    *storage.Storage
	logger     *logging.Logger
	server     *server.Server
}

func New() *App {
	app := &App{}

	loggerCfg, err := getLoggerConfig()
	if err != nil {
		panic("Can`t get logger config. Error: " + err.Error())
	}

	logger := logging.NewLogger(loggerCfg)

	app.logger = logger

	ctx := context.TODO()

	if err := godotenv.Load(".env"); err != nil {
		logger.Error("Can`t load env. Error: " + err.Error())
	}

	postgresCfg, err := getPostgresConfig()
	if err != nil {
		logger.Error("Can`t get postgres config. Error: " + err.Error())
	}

	pgPool, err := postgresql.ConnectToDB(ctx, postgresCfg)
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

	redisCfg, err := getRedisConfig()
	if err != nil {
		logger.Error("Can`t get redis config. Error: " + err.Error())
	}

	redisClient, err := redis.ConnectToRedis(ctx, redisCfg)
	if err != nil {
		logger.Error("Can`t connect to redis. Error: " + err.Error())
	} else {
		logger.Info("Redis is connected")
	}

	serverCfg, err := getServerConfig()
	if err != nil {
		logger.Error("Can`t get server config. Error: " + err.Error())
	}

	app.storage = storage.New(pgPool, redisClient)

	app.usecase = usecase.New(app.storage)

	app.controller = controller.New(app.usecase)

	handler := app.controller.InitRoutes()

	app.server = server.New(handler, serverCfg)

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