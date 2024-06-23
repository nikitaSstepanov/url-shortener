package server

import (
	"net/http"
	"log/slog"
	"context"
	"time"
	"fmt"
)

type Server struct {
	server *http.Server
}

type Config struct {
	Url string `yaml:"url"`
}

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 5 * time.Second
)

func New(handler http.Handler, cfg *Config) *Server {
	return &Server{
		server: &http.Server{
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			Handler:      handler,
			Addr:         cfg.Url,
		},
	}
}

func (s *Server) Start(logger *slog.Logger) error {
	var err error

	logger.Info(fmt.Sprintf("Server started at %s", s.server.Addr))

	go func() {
		err = s.server.ListenAndServe()
	}()

	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}