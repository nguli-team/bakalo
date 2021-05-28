package persistence

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"bakalo.li/internal/config"
	"bakalo.li/internal/logger"
)

const (
	maxDBConnRetry = 3
	retryDelay     = 5
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

	gormConfig := &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	}

	var db *gorm.DB
	var err error

	// if error, retry to connect to database
	for ret := 0; ret < maxDBConnRetry; ret++ {
		db, err = gorm.Open(postgres.Open(dsn), gormConfig)
		if err != nil {
			logger.Log().Errorf(
				"failed to initialize database, retrying in %ds [%d/%d]: %v",
				retryDelay,
				ret+1,
				maxDBConnRetry,
				err,
			)
			time.Sleep(retryDelay * time.Second)
		}
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
