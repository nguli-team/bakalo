package http

import (
	bakaloMiddleware "bakalo.li/internal/application/http/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

func InitChiRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(bakaloMiddleware.RequestLogger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)

		_, _ = w.Write([]byte("welcome"))
	})

	return r
}
