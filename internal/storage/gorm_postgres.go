package storage

import (
	"bakalo.li/internal/config"
	"bakalo.li/internal/logger"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

// NewGormPostgres creates a new storage connection with gorm
func NewGormPostgres(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=%s",
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.User,
		cfg.Password,
		cfg.TimeZone,
	)

	gormConfig := &gorm.Config{}

	// configure zap as logger if logger is of type *zap.SugaredLogger
	zapSLogger, ok := logger.Log.(*zap.SugaredLogger)
	if ok {
		gzl := zapgorm2.New(zapSLogger.Desugar())
		gzl.SetAsDefault()
		gormConfig.Logger = gzl
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
