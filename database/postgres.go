package database

import (
	"fmt"
	"github.com/nguli-team/bakalo/config"
	"github.com/nguli-team/bakalo/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresDatabase creates a new database connection with gorm
func NewPostgresDatabase(cfg config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=%s",
		cfg.Hostname,
		cfg.Port,
		cfg.Database,
		cfg.User,
		cfg.Password,
		cfg.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	exception.PanicIfNeeded(err)

	return db
}
