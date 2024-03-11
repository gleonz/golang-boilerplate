package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/gleonz/golang-boilerplate/config"
	v1 "github.com/gleonz/golang-boilerplate/internal/controller/http/v1"
	"github.com/gleonz/golang-boilerplate/pkg/httpserver"
	"github.com/gleonz/golang-boilerplate/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Here you would initialize the new business logic,
	// after you have modified the usecase package according to your needs.
	// For example:
	// myNewUseCase := NewMyUseCase()

	// HTTP Server
	handler := gin.New()
	// Assuming v1.NewRouter and the business logic adjust based on your changes.
	v1.NewRouter(handler, l /*, myNewUseCase*/)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Espera de señal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	if err := httpServer.Shutdown(); err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
