package cmd

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"

	bakaloHttp "github.com/nguli-team/bakalo/internal/application/http"
)

func newServeCmd() *cobra.Command {
	serveCmd := &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			ctx, cancel := context.WithCancel(context.Background())
			go func() {
				<-c
				cancel()
			}()
			bakaloHttp.Serve(
				ctx, &bakaloHttp.ServeConfig{
					Config:              cfg,
					RequestLoggerOutput: os.Stdout,
				},
			)
		},
	}
	return serveCmd
}
