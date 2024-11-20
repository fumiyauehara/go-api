package middleware

import (
	"gorm.io/gorm"
	"net/http"
)

func SetDBSessionVariable(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reader, ok := r.Context().Value(DBConnCtxKey).(*gorm.DB)
		if !ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		tenantId, ok := r.Context().Value(TenantIdCtxKey).(int)
		if !ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		reader.Exec("SET @target_tenant_id = ?;", tenantId)

		next.ServeHTTP(w, r)
	})
}
