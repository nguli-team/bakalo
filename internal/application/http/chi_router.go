package http

import (
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"

	"bakalo.li/internal/application/http/handler"
	bakaloMiddleware "bakalo.li/internal/application/http/middleware"
	"bakalo.li/internal/config"
	"bakalo.li/internal/logger"
)

func NewChiRouter(
	env config.Environment,
	loggerOutput io.Writer,
	services *ServiceContainer,
) *chi.Mux {
	router := chi.NewRouter()

	// middlewares
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
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

	// handlers
	boardHandler := handler.NewBoardHandler(services.BoardService)
	threadHandler := handler.NewThreadHandler(services.ThreadService)

	router.Route(
		"/v1", func(r chi.Router) {
			// board endpoints
			r.Get("/boards", boardHandler.ListBoards)
			r.Get("/board/{id:[0-9]+}", boardHandler.GetByID)
			r.Get("/board/{shorthand:[a-z]+}", boardHandler.GetByShorthand)

			// thread endpoints
			r.Get("/threads", threadHandler.ListThreads)
			r.Get("/thread/{id:[0-9]+}", threadHandler.GetByID)
			r.Post("/thread", threadHandler.CreateThreadMultipart)
		},
	)

	router.Get(
		"/ping", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("pong"))
		},
	)

	return router
}
