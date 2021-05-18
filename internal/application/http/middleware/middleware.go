package middleware

import (
	"bakalo.li/internal/logger"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

func RequestLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		t1 := time.Now()
		defer func() {
			logger.Log.Infof(
				"HTTP request: %s %s %s %d %dbytes %s %dms",
				r.RemoteAddr,
				r.Method,
				r.URL,
				ww.Status(),
				ww.BytesWritten(),
				ww.Header(),
				time.Since(t1).Milliseconds(),
			)
		}()
		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}
