package http

import (
	"bakalo.li/internal/config"
	"bakalo.li/internal/logger"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ServerConfig struct {
	Env                 config.Environment
	RequestLoggerOutput io.Writer
	HTTPServerConfig    config.HTTPServerConfig
}

func Serve(ctx context.Context, cfg *ServerConfig) {
	services, err := NewServiceContainer(cfg.DatabaseConfig)
	if err != nil {
		logger.Log.Fatal(err)
	}

	router := NewChiRouter(env, services)

	addr := fmt.Sprintf("%s:%d", srvCfg.Hostname, srvCfg.Port)
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

	// put clean up code here

	err = srv.Shutdown(ctxShutdown)
	if err != nil {
		logger.Log.Fatal(err)
	}
	logger.Log.Info("http server stopped properly")
}
