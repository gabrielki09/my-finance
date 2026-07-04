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

func runUp(ctx context.Context, db *pgxpool.Pool) (err error) {
	if _, err := db.Exec(
		ctx,
		upSchemaSQL,
	); err != nil {

		return err
	}

	return err
}

func runDown(ctx context.Context, db *pgxpool.Pool) (err error) {
	if _, err := db.Exec(
		ctx,
		downSchemaSQL,
	); err != nil {

		return err
	}

	return err
}

func HandleMigrate(ctx context.Context, db *pgxpool.Pool, migrateType string) (err error) {
	logger.General.Info.Println("migrateType: ", migrateType)

	if migrateType == "migrate" {
		if err := runUp(ctx, db); err != nil {
			logger.General.Error.Println("Erro ao roda a migrate:", err)
			return err
		}

		logger.General.Info.Println("Novas tabelas criadas com sucesso!")
	}

	if migrateType == "migrate:fre" {
		if err := runDown(ctx, db); err != nil {
			logger.General.Error.Println("Erro ao dropar todas as tabelas:", err)
			return err
		}
		logger.General.Info.Println("Tabelas dropadas com sucesso!")

		if err := runUp(ctx, db); err != nil {
			logger.General.Error.Println("Erro ao roda a migrate:", err)
			return err
		}

		logger.General.Info.Println("Tabelas recriadas com sucesso!")
	}

	return err
}
