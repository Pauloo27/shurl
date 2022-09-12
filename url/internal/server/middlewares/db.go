package middlewares

import (
	"context"
	"net/http"

	"github.com/Pauloo27/shurl/url/internal/service"
)

type CtxKey string

const (
	DBCtxKey = CtxKey("db")
)

func NewDBCtx(service *service.URLService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), DBCtxKey, service.DB)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
