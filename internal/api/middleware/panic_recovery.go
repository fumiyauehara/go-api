package middleware

import (
	"net/http"
	"runtime/debug"
)

// RecoverOccurredPanicFromGoRoutine 特定のgoroutineでpanicが起きるとアプリ自体がクラッシュするのでアプリ全体には影響を与えないようにrecoveryする。
func RecoverOccurredPanicFromGoRoutine(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recovered := recover(); recovered != nil {
				debug.PrintStack()
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
