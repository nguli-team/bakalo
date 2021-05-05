package main

import (
	"github.com/nguli-team/bakalo/config"
	"github.com/nguli-team/bakalo/database"
	"github.com/nguli-team/bakalo/logger"
	"os"
)

func main() {
	// fetch configurations from file
	cfgFile := os.Getenv("CONFIG_FILE")
	if cfgFile == "" {
		// default if config file is not specified in env var
		cfgFile = "config.yml"
	}
	cfg, err := config.NewConfig(cfgFile)
	if err != nil {
		panic(err)
	}

	// logger configuration
	zapLogger := logger.NewZapSugaredLogger(cfg.Env)
	logger.SetLogger(zapLogger)

	// connect to database
	dbCfg := cfg.Server.Database
	dbConn := database.NewPostgresDatabase(dbCfg)

	// perform database migration if specified
	if dbCfg.AutoMigrate {
		// TODO: Add models for migrations
		err := dbConn.AutoMigrate()
		if err != nil {
			logger.Log.Fatal("Database migration failed!")
		}
		logger.Log.Info("Database auto-migration completed.")
	}
}
