package http

import (
	"io"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"

	"github.com/nguli-team/bakalo/internal/application/http/handler"
	"github.com/nguli-team/bakalo/internal/application/http/middleware"
	"github.com/nguli-team/bakalo/internal/config"
	"github.com/nguli-team/bakalo/internal/logger"
)

func NewChiRouter(
	env config.Environment,
	loggerOutput io.Writer,
	services *ServiceContainer,
) *chi.Mux {
	router := chi.NewRouter()

	// middlewares
	initMiddlewares(router, env, loggerOutput)

	// handlers
	boardHandler := handler.NewBoardHandler(services.BoardService)
	threadHandler := handler.NewThreadHandler(services.ThreadService)
	postHandler := handler.NewPostHandler(services.PostService)
	vipHandler := handler.NewTokenHandler(services.TokenService)

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

			// post endpoints
			r.Get("/posts", postHandler.ListPosts)
			r.Post("/post", postHandler.CreatePostMultipart)
			r.Delete("/post/{id:[0-9]+}", postHandler.DeletePost)

			r.Post("/vip", vipHandler.RequestNewToken)
			r.Post("/vip/login", vipHandler.Login)
		},
	)

	return router
}

func initMiddlewares(router *chi.Mux, env config.Environment, loggerOutput io.Writer) {
	router.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{"https://*", "http://*"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: true,
				MaxAge:           300,
			},
		),
	)

	router.Use(chiMiddleware.RealIP)
	router.Use(middleware.RequestIP)
	router.Use(chiMiddleware.RequestID)

	// logger middleware
	if env == config.Production {
		reqLogger := logrus.New()
		reqLogger.Formatter = &logrus.JSONFormatter{
			DisableTimestamp: true,
		}
		reqLogger.Out = loggerOutput
		router.Use(middleware.NewStructuredLogger(reqLogger))
	} else {
		router.Use(
			chiMiddleware.RequestLogger(&chiMiddleware.DefaultLogFormatter{Logger: logger.Log()}),
		)
	}

	router.Use(chiMiddleware.Recoverer)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(chiMiddleware.Heartbeat("/ping"))
}
