package middleware

import (
	"context"
	"github.com/fumiyauehara/go-api/internal/api/util"
	"net/http"
)

const TenantIdCtxKey = "tenant-id"

func SetTenantId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenantId := util.ConvertStringToInt(r.Header.Get("X-Tenant-Id"))
		if tenantId == 0 {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), TenantIdCtxKey, tenantId)))
	})
}
