package api

import (
	"github.com/fumiyauehara/go-api/internal/api/handler"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", handler.IndexHandler).Methods("GET")
	api.HandleFunc("/index", handler.IndexHandler).Methods("GET")

	sse := r.PathPrefix("/sse").Subrouter()
	sse.HandleFunc("/", handler.EventIndexHandler).Methods("GET")
	sse.HandleFunc("/index", handler.EventIndexHandler).Methods("GET")

	return r
}
