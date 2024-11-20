package middleware

import (
	"context"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

const (
	DBConnCtxKey = "db-conn"
)

func MakeSettingDbConnMiddleware(dbConn *gorm.DB) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), DBConnCtxKey, dbConn.Session(&gorm.Session{}))

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
