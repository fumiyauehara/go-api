package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func RecoverOccurredPanicOnSseGoroutine(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, channel := context.WithCancel(r.Context())
		defer channel()

		errChan := make(chan error, 1)
		go func() {
			defer func() {
				if recovered := recover(); recovered != nil {
					log.Printf("panic in SSE handler: %v\n", recovered)
					debug.PrintStack()
					errChan <- fmt.Errorf("panic in SSE handler: %v", recovered)
				}
			}()
			next.ServeHTTP(w, r.WithContext(ctx))
			errChan <- nil
		}()

		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
		case err := <-errChan:
			if err != nil {
				log.Println("Error in SSE handler:", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
	})
}
