package server

import (
	"fmt"
	"net/http"

	"github.com/Pauloo27/shurl/url/internal/service"
	"github.com/Pauloo27/shurl/url/internal/server/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func Start(logger *zap.SugaredLogger, service *service.URLService) error {
	logger.Infof("Starting HTTP server at port %d...", service.Http.Port)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middlewares.NewZapMiddleware("router", logger))
	r.Use(middleware.Recoverer)

	return http.ListenAndServe(fmt.Sprintf(":%d", service.Http.Port), r)
}
