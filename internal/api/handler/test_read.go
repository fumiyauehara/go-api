package handler

import (
	"encoding/json"
	"github.com/fumiyauehara/go-api/internal/api/middleware"
	"github.com/fumiyauehara/go-api/internal/api/model"
	"gorm.io/gorm"
	"net/http"
)

func Read(w http.ResponseWriter, r *http.Request) {
	dbConn, ok := r.Context().Value(middleware.DBConnCtxKey).(*gorm.DB)
	if !ok {
		http.Error(w, "db conn not found", http.StatusInternalServerError)
		return
	}

	var results []model.ViewEmployee
	if err := dbConn.Find(&results).Error; err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	if len(results) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		return
	}
}
