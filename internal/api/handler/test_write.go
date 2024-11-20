package handler

import (
	"encoding/json"
	"fmt"
	"github.com/fumiyauehara/go-api/internal/api/middleware"
	"github.com/fumiyauehara/go-api/internal/api/model"
	"gorm.io/gorm"
	"net/http"
)

func Write(w http.ResponseWriter, r *http.Request) {
	dbConn, ok := r.Context().Value(middleware.DBConnCtxKey).(*gorm.DB)
	if !ok {
		http.Error(w, "db conn not found in ctx", http.StatusInternalServerError)
		return
	}
	tenantId, ok := r.Context().Value(middleware.TenantIdCtxKey).(int)
	if !ok {
		http.Error(w, "tenant not found in ctx", http.StatusInternalServerError)
		return
	}

	var reqEmployees []model.RequestEmployee
	err := json.NewDecoder(r.Body).Decode(&reqEmployees)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed parsing json. reason: %s\n", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var employees []model.Employee
	for _, reqEmp := range reqEmployees {
		employees = append(employees, model.Employee{
			Name:     reqEmp.Name,
			Email:    reqEmp.Email,
			TenantID: tenantId,
		})
	}

	result := dbConn.Create(&employees)
	if result.Error != nil {
		http.Error(w, "db create employee error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
