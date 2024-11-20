package middleware

import (
	"github.com/fumiyauehara/go-api/internal/api/model"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

// MakeTenantIdValidator reader専用のview databaseはテナントIDがあることを前提んしているため、
//
//	ここではテナントID書き込み専用のDBへのconnectionを渡している。なお、命名で明示する必要はないが、
//	分かりやすさのためにそうしているが、実際には呼び出し元次第である。
func MakeTenantIdValidator(dbConn *gorm.DB) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tenantId, ok := r.Context().Value(TenantIdCtxKey).(int)
			if !ok {
				http.Error(w, "tenant id not found in ctx", http.StatusNotFound)
				return
			}

			var cnt int64
			dbConn.Model(&model.Tenant{}).Where("id = ?", tenantId).Count(&cnt)

			if cnt == 0 {
				http.Error(w, "tenant not found", http.StatusNotFound)
				return
			} else if cnt >= 2 {
				http.Error(w, "tenant More than one record found", http.StatusConflict)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
