package main

import (
	"log"

	"github.com/cafasaru/nats_starter/config"
	"github.com/cafasaru/nats_starter/pkg/logger"
)

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Info("Starting emails microservice")
	appLogger.Infof(
		"AppVersion: %s, LogLevel: %s, DevelopmentMode: %t",
		cfg.AppVersion,
		cfg.Logger.Level,
		cfg.HTTP.Development,
	)
	appLogger.Infof("Success loaded config: %+v", cfg.AppVersion)

}
