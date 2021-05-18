package cmd

import (
	"bakalo.li/internal/config"
	"bakalo.li/internal/domain"
	"bakalo.li/internal/logger"
	"bakalo.li/internal/repository"
	"bakalo.li/internal/storage"
	"context"
	"github.com/spf13/cobra"
)

var tableOnly bool

func newMigrateCmd() *cobra.Command {
	migrateCmd := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			err := migrateGorm(cfg, tableOnly)
			if err != nil {
				logger.Log.Fatal(err)
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
	db, err := storage.NewGormPostgres(cfg.Server.Database)
	if err != nil {
		return err
	}

	if tableOnly {
		logger.Log.Info("starting table only migration")
	} else {
		logger.Log.Info("starting migration")
	}

	// migrate tables
	err = db.AutoMigrate(
		&domain.Board{},
	)
	if err != nil {
		return err
	}

	// return early if only creating tables
	if tableOnly {
		return nil
	}

	// create boards from configuration
	boardRepo := repository.NewGormBoardRepository(db)
	for _, b := range cfg.App.Boards {
		ctx := context.TODO()
		err = boardRepo.Create(ctx, &domain.Board{
			Shorthand:   b.Shorthand,
			Name:        b.Name,
			Description: b.Description,
		})
		if err != nil {
			return err
		}
	}

	logger.Log.Info("[MIGRATION] migration successful")
	return nil
}
