package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"bakalo.li/internal/config"
	"bakalo.li/internal/logger"
)

type ServeConfig struct {
	Config              config.Config
	RequestLoggerOutput io.Writer
}

func Serve(ctx context.Context, cfg *ServeConfig) {
	services, err := NewServiceContainer(cfg.Config.Database)
	if err != nil {
		logger.Log().Fatal(err)
	}

	router := NewChiRouter(cfg.Config.Env, cfg.RequestLoggerOutput, services)

	serverCfg := cfg.Config.Server
	addr := fmt.Sprintf("%s:%d", serverCfg.Hostname, serverCfg.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Log().Fatal(err)
		}
	}()
	logger.Log().Info("http server is listening at ", addr)

	<-ctx.Done()

	logger.Log().Info("stopping http server")
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// put clean up code here

	err = srv.Shutdown(ctxShutdown)
	if err != nil {
		logger.Log().Fatal(err)
	}
	logger.Log().Info("http server stopped properly")
}
