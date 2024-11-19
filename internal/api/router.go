package api

import (
	"github.com/fumiyauehara/go-api/internal/api/handler"
	"github.com/fumiyauehara/go-api/internal/api/middleware"
	"github.com/fumiyauehara/go-api/internal/api/model"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouter(conn model.DBConn) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.RecoverOccurredPanicFromGoroutine)
	api.Use(middleware.ValidateTenantID)
	api.Use(middleware.MakeSettingContextMiddleware(conn))
	api.Use(middleware.SetDBSessionVariable)
	api.HandleFunc("/", handler.IndexHandler).Methods("GET")
	api.HandleFunc("/index", handler.IndexHandler).Methods("GET")
	api.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("recovery confirm")
	})
	api.HandleFunc("/reader-conn", handler.ReaderConn).Methods("GET")

	sse := r.PathPrefix("/sse").Subrouter()
	sse.Use(middleware.RecoverOccurredPanicFromGoroutine)
	sse.Use(middleware.RecoverOccurredPanicOnSseGoroutine)
	sse.Use(middleware.SetSseHeader)
	sse.HandleFunc("/", handler.EventIndexHandler).Methods("GET")
	sse.HandleFunc("/index", handler.EventIndexHandler).Methods("GET")

	return r
}
