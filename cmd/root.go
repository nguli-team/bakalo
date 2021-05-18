package cmd

import (
	"bakalo.li/internal/config"
	"bakalo.li/internal/logger"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile string
	cfg     config.Config
)

func init() {
	cfgFile = os.Getenv("CONFIG_FILE")
	if cfgFile == "" {
		cfgFile = "config.yml"
	}
	var err error
	cfg, err = config.NewConfig(cfgFile)
	if err != nil {
		panic(err)
	}

	// set global logger
	l, err := logger.NewZapSugaredLogger(cfg.Env)
	if err != nil {
		panic(err)
	}
	logger.SetLogger(l)
}

// Execute executes the root command.
func Execute() error {
	migrateCmd := newMigrateCmd()
	serveCmd := newServeCmd()

	rootCmd := &cobra.Command{
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger.Log.Info("using configuration file: ", cfgFile)
			logger.Log.Info("running in environment: ", cfg.Env)
		},
	}
	rootCmd.AddCommand(migrateCmd, serveCmd)
	return rootCmd.Execute()
}
