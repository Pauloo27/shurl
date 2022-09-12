package routes

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"
)

func RouteRedirect() http.Handler {
	r := chi.NewRouter()

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		path, err := parsePath(r.URL)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		if path == "" {
			http.Redirect(w, r, "/_/version", http.StatusTemporaryRedirect)
		}

		// TODO: get the URL from the database
		// TODO: redirect =)

		http.NotFound(w, r)
	})
	return r
}

func parsePath(url *url.URL) (string, error) {
	rawPath := url.Path
	pathParts := strings.Split(strings.Trim(rawPath, "/"), "/")
	if len(pathParts) != 1 {
		return "", errors.New("invalid path")
	}
	return strings.TrimSpace(pathParts[0]), nil
}
