package migration

import (
	"context"
	_ "embed"
	"finance/internal/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed sql/up.sql
var upSchemaSQL string

//go:embed sql/down.sql
var downSchemaSQL string

func RunMigrate(db *pgxpool.Pool) (err error) {
	ctx := context.Background()

	logger.General.Info.Println("Script:", upSchemaSQL)

	tx, err := db.Begin(ctx)

	if err != nil {
		return err
	}

	if _, err := tx.Exec(
		ctx,
		upSchemaSQL,
	); err != nil {
		logger.General.Error.Println("Erro ao rodar a migrate:", err)
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		logger.General.Error.Println("Erro ao commitar a migrate:", err)
		return err
	}

	return err
}
