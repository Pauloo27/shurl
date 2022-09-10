package server

import (
	"net/http"

	"github.com/Pauloo27/shurl/url/internal/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Start() error {
	logger.L.Info("Starting HTTP server...")

	r := chi.NewRouter()

	// TODO: logger middleware

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return http.ListenAndServe(":3000", r)
}
