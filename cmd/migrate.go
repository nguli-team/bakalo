package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/nguli-team/bakalo/internal/config"
	"github.com/nguli-team/bakalo/internal/domain"
	"github.com/nguli-team/bakalo/internal/logger"
	"github.com/nguli-team/bakalo/internal/repository"
	"github.com/nguli-team/bakalo/internal/storage/persistence"
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
	db, err := persistence.NewGormPostgres(cfg.Database)
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
		&domain.VipToken{},
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
			logger.Log().Error(err)
		}
	}

	if err != nil {
		logger.Log().Warn("migration finished with error")
		return nil
	}

	logger.Log().Info("migration successful")
	return nil
}
