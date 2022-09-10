package bootstrap

import (
	"github.com/Pauloo27/shurl/url/internal/logger"
	"github.com/Pauloo27/shurl/url/internal/server"
)

func handleFatal(err error) {
	if err != nil {
		logger.L.Fatal(err)
	}
}

func Start() {
	logger.L.Info("Starting URL service...")
	handleFatal(server.Start())
}
