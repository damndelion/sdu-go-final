package main

import (
	"github.com/damndelion/sdu-go-final/config"
	"github.com/damndelion/sdu-go-final/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	// Configuration test
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
