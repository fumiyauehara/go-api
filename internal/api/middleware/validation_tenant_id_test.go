package middleware_test

import (
	"context"
	"github.com/fumiyauehara/go-api/internal/api/middleware"
	"github.com/fumiyauehara/go-api/internal/api/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidateTenantId(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	tx.Create(&model.Tenant{
		ID:      10,
		Name:    "test",
		Address: "test@example.com",
		Tel:     "111-222-333",
	})

	// モックの次のハンドラー
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	})

	// テスト用のリクエストとレスポンス
	tests := []struct {
		name           string
		tenantId       interface{}
		dbConn         interface{}
		expectedStatus int
	}{
		{
			name:           "Valid Tenant ID",
			tenantId:       10,
			dbConn:         tx,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid Tenant ID",
			tenantId:       999,
			dbConn:         tx,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "No Tenant ID in Context",
			tenantId:       nil,
			dbConn:         tx,
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// リクエスト作成
			req := httptest.NewRequest(http.MethodGet, "/test", nil)

			// コンテキストに必要なデータを設定
			ctx := req.Context()
			if tt.tenantId != nil {
				ctx = context.WithValue(ctx, middleware.TenantIdCtxKey, tt.tenantId)
			}
			req = req.WithContext(ctx)

			// レスポンスレコーダー
			rr := httptest.NewRecorder()

			// ハンドラー呼び出し
			targetMiddleware := middleware.MakeTenantIdValidator(tx)
			handler := targetMiddleware(nextHandler)
			handler.ServeHTTP(rr, req)

			// ステータスコードの検証
			assert.Equal(t, tt.expectedStatus, rr.Code)
		})
	}
}
