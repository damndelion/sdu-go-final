// Package app configures and runs application.
package app

import (
	"fmt"
	v1 "github.com/evrone/go-clean-template/internal/controller/http"
	"github.com/evrone/go-clean-template/internal/entity"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/internal/usecase/repo"
	"github.com/evrone/go-clean-template/pkg/httpserver"
	"github.com/evrone/go-clean-template/pkg/postgres"
	"github.com/sirupsen/logrus"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {

	log := logrus.New()
	// Repository
	db, err := postgres.New(cfg.PG.URL)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer postgres.Close(db)

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Error("failed to migrate")
		return
	}

	// Use case
	userUseCase := usecase.New(
		repo.New(db),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, log, userUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
