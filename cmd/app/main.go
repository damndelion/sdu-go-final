package main

import (
	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
