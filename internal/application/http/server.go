package http

import (
	"bakalo.li/internal/config"
	"bakalo.li/internal/logger"
	"context"
	"fmt"
	"net/http"
	"time"
)

func Serve(ctx context.Context, cfg config.ServerConfig) {
	router := InitChiRouter()

	addr := fmt.Sprintf("%s:%d", cfg.Hostname, cfg.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal(err)
		}
	}()
	logger.Log.Info("http server is listening at ", addr)

	<-ctx.Done()

	logger.Log.Info("stopping http server")
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := srv.Shutdown(ctxShutdown)
	if err != nil {
		logger.Log.Fatal(err)
	}
	logger.Log.Info("http server stopped properly")
}
