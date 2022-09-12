package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RouteUnderscore() http.Handler {
	r := chi.NewRouter()
	r.Get("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("SHURL v0.0.1 PRE-ALPHA"))
	})
	return r
}
