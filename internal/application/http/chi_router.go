package http

import (
	"bakalo.li/internal/application/http/handler"
	bakaloMiddleware "bakalo.li/internal/application/http/middleware"
	"bakalo.li/internal/config"
	"bakalo.li/internal/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func NewChiRouter(env config.Environment, loggerOutput io.Writer, services ServiceContainer) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)

	if env == config.Production {
		reqLogger := logrus.New()
		reqLogger.Formatter = &logrus.JSONFormatter{
			DisableTimestamp: true,
		}
		reqLogger.Out = loggerOutput
		router.Use(bakaloMiddleware.NewStructuredLogger(reqLogger))
	} else {
		router.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger.Log}))
	}

	router.Use(middleware.Recoverer)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	boardHandler := handler.NewBoardHandler(services.BoardService)

	router.Route("/v1", func(r chi.Router) {
		r.Get("/boards", boardHandler.ListBoards)
		r.Get("/board/{id}", boardHandler.GetByID)
	})

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	return router
}
