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
	pg, err := db.Connect(config)
	handleFatal(sugar, err)

	err = pg.Ping()
	handleFatal(sugar, err)

	applied, err := db.Migrate(pg)
	handleFatal(sugar, err)

	sugar.Infof("Applied %d migrations", applied)

	service := service.NewService(config, pg)

	handleFatal(sugar, server.Start(sugar, service))
}
