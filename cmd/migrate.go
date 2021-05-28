package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"bakalo.li/internal/config"
	"bakalo.li/internal/domain"
	"bakalo.li/internal/logger"
	"bakalo.li/internal/repository"
	"bakalo.li/internal/storage"
)

var tableOnly bool

func newMigrateCmd() *cobra.Command {
	migrateCmd := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			err := migrateGorm(cfg, tableOnly)
			if err != nil {
				logger.Log().Fatal(err)
			}
		},
	}
	migrateCmd.Flags().BoolVarP(
		&tableOnly,
		"create-table-only",
		"",
		false,
		"only create table(s) and skip data seeding.",
	)
	return migrateCmd
}

func migrateGorm(cfg config.Config, tableOnly bool) error {
	db, err := storage.NewGormPostgres(cfg.Database)
	if err != nil {
		return err
	}

	if tableOnly {
		logger.Log().Info("starting table only migration")
	} else {
		logger.Log().Info("starting migration")
	}

	// migrate tables
	err = db.AutoMigrate(
		&domain.Board{},
		&domain.Thread{},
		&domain.Post{},
	)
	if err != nil {
		return err
	}

	// return early if only creating tables
	if tableOnly {
		logger.Log().Info("table migration successful")
		return nil
	}

	// create boards from configuration
	boardRepo := repository.NewGormBoardRepository(db)
	for _, b := range cfg.App.Boards {
		ctx := context.TODO()
		_, err = boardRepo.Create(
			ctx, &domain.Board{
				Shorthand:   b.Shorthand,
				Name:        b.Name,
				Description: b.Description,
			},
		)
		if err != nil {
			return err
		}
	}

	logger.Log().Info("migration successful")
	return nil
}
