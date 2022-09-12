package routes

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Pauloo27/shurl/url/internal/server/middlewares"
	"github.com/Pauloo27/shurl/url/internal/service"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

const (
	GenericDatabaseError = "Something went wrong while reaching the database..."
)

func RouteRedirect(logger *zap.SugaredLogger) http.Handler {
	r := chi.NewRouter()

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		path, err := parsePath(r.URL)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		if path == "" {
			http.Redirect(w, r, "/_/version", http.StatusTemporaryRedirect)
			return
		}

		service := r.Context().Value(middlewares.ServiceCtxKey).(*service.URLService)
		db := service.DB

		res, err := db.Query("SELECT long_url FROM url WHERE path = $1", path)
		if err != nil {
			http.Error(w, GenericDatabaseError, http.StatusInternalServerError)
			logger.Error(err)
			return
		}
		defer res.Close()

		var longURL string
		if !res.Next() {
			http.NotFound(w, r)
			return
		}
		err = res.Scan(&longURL)
		if err != nil {
			http.Error(w, GenericDatabaseError, http.StatusInternalServerError)
			logger.Error(err)
		}

		http.Redirect(w, r, longURL, http.StatusTemporaryRedirect)
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
