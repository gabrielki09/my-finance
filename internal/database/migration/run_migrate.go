package migration

import (
	_ "embed"
	"finance/internal/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed schema.sql
var schemaSQL string

func RunMigrate(db *pgxpool.Pool) (err error) {
	logger.General.Info.Println("Script:", schemaSQL)
	return err
}
