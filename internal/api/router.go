package api

import (
	"github.com/fumiyauehara/go-api/internal/api/handler"
	"github.com/fumiyauehara/go-api/internal/api/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.RecoverOccurredPanicFromGoRoutine)
	api.HandleFunc("/", handler.IndexHandler).Methods("GET")
	api.HandleFunc("/index", handler.IndexHandler).Methods("GET")
	api.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("recovery confirm")
	})

	sse := r.PathPrefix("/sse").Subrouter()
	sse.Use(middleware.RecoverOccurredPanicFromGoRoutine)
	sse.Use(middleware.SetSseHeader)
	sse.HandleFunc("/", handler.EventIndexHandler).Methods("GET")
	sse.HandleFunc("/index", handler.EventIndexHandler).Methods("GET")

	return r
}
