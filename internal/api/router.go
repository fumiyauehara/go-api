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
	api.Use(middleware.SetTenantId)

	immutableApi := api.Methods("GET").Subrouter()
	immutableApi.Use(middleware.MakeSettingDbConnMiddleware(conn.Reader))
	immutableApi.Use(middleware.ValidateTenantId)
	immutableApi.Use(middleware.SetDBSessionVariable)
	immutableApi.HandleFunc("/", handler.IndexHandler).Methods("GET")
	immutableApi.HandleFunc("/index", handler.IndexHandler).Methods("GET")
	immutableApi.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("recovery confirm")
	})
	immutableApi.HandleFunc("/read", handler.Read).Methods("GET")

	mutableApi := api.Methods("POST", "PATCH", "PUT", "DELETE").Subrouter()
	mutableApi.Use(middleware.MakeSettingDbConnMiddleware(conn.Writer))
	immutableApi.Use(middleware.ValidateTenantId)
	mutableApi.HandleFunc("/write", handler.Write).Methods("POST")

	sse := r.PathPrefix("/sse").Subrouter()
	sse.Use(middleware.RecoverOccurredPanicFromGoroutine)
	sse.Use(middleware.RecoverOccurredPanicOnSseGoroutine)
	sse.Use(middleware.SetSseHeader)
	sse.HandleFunc("/", handler.EventIndexHandler).Methods("GET")
	sse.HandleFunc("/index", handler.EventIndexHandler).Methods("GET")

	return r
}
