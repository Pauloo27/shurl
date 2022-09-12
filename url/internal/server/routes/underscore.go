package routes

import (
	"net/http"

	"github.com/Pauloo27/shurl/url/internal/server/middlewares"
	"github.com/Pauloo27/shurl/url/internal/service"
	"github.com/go-chi/chi/v5"
)

func RouteUnderscore() http.Handler {
	r := chi.NewRouter()
	r.Get("/version", func(w http.ResponseWriter, r *http.Request) {
		service := r.Context().Value(middlewares.ServiceCtxKey).(*service.URLService)
		w.Write([]byte("SHURL " + service.Version))
	})
	return r
}
