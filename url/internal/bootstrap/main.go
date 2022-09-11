package bootstrap

import (
	"github.com/Pauloo27/shurl/url/internal/db"
	"github.com/Pauloo27/shurl/url/internal/server"
	"github.com/Pauloo27/shurl/url/internal/service"
	"go.uber.org/zap"
)

func handleFatal(logger *zap.SugaredLogger, err error) {
	if err != nil {
		logger.Fatal(err)
	}
}

func Start() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	sugar := logger.Sugar()

	sugar.Info("Starting URL service...")

	sugar.Infof("Loading configuration from file %s...", ConfigFileName)
	config, err := LoadConfig()
	handleFatal(sugar, err)

	sugar.Info("Connecting to Postgres...")
	db, err := db.Connect(config)
	handleFatal(sugar, err)

	err = db.Ping()
	handleFatal(sugar, err)

	service := service.NewService(config, db)

	handleFatal(sugar, server.Start(sugar, service))
}
