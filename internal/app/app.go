// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/damndelion/sdu-go-final/config"
	"github.com/damndelion/sdu-go-final/internal/controller/http"
	"github.com/damndelion/sdu-go-final/internal/entity"
	"github.com/damndelion/sdu-go-final/internal/usecase"
	"github.com/damndelion/sdu-go-final/internal/usecase/repo"
	"github.com/damndelion/sdu-go-final/pkg/httpserver"
	"github.com/damndelion/sdu-go-final/pkg/postgres"

	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

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

	err = db.AutoMigrate(&entity.User{}, &entity.Menu{}, &entity.Order{})
	if err != nil {
		log.Error("failed to migrate")
		return
	}

	// Use case
	userUseCase := usecase.NewUserUseCase(
		repo.NewUserRepo(db),
	)
	authUseCase := usecase.NewAuthUseCase(
		cfg,
		repo.NewUserRepo(db),
	)
	menuUseCase := usecase.NewMenuUseCase(
		repo.NewMenuRepo(db),
	)

	orderUseCase := usecase.NewOrderUseCase(
		repo.NewOderRepo(db),
	)

	// HTTP Server
	handler := gin.New()
	http.NewRouter(handler, log, orderUseCase, userUseCase, authUseCase, menuUseCase, cfg)
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
