package db

import (
	"context"
	"finance/internal/config"
	"finance/internal/database"
	"finance/internal/logger"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	upContentToRun   string
	downContentToRun string
)

func connDB() (*pgxpool.Pool, error) {
	configDB, err := config.LoadDbConfig()
	if err != nil {
		return nil, err
	}

	conn, err := database.ConnectDb(configDB)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func readContentToRun(rootPath string, upDB, downDB bool) error {
	contentFiles := filepath.Join(rootPath, "internal", "database", "migration", "sql")

	if upDB {
		upSql, err := os.ReadFile(contentFiles + "/up.sql")
		if err != nil {
			logger.Error("Erro ao ler o arquivo up.sql", err)
			return err
		}

		upContentToRun = string(upSql)
		return nil
	}

	if downDB {
		downSql, err := os.ReadFile(contentFiles + "/down.sql")
		if err != nil {
			logger.Error("Erro ao ler o arquivo down.sql", err)
			return err
		}

		downContentToRun = string(downSql)
		return nil
	}

	return nil
}

func backupDB(ctx context.Context, rootPath string) (string, error) {
	backupDir := filepath.Join(rootPath, "backup")

	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return "", fmt.Errorf("erro ao criar diretório de backups: %w", err)
	}

	fileNameDump := fmt.Sprintf(
		"finance_%s.dump",
		time.Now().Format("20060102_150405"),
	)

	backupPath := filepath.Join(backupDir, fileNameDump)

	cmd := exec.CommandContext(
		ctx,
		"pg_dump",
		"-h", "postgres",
		"-p", "5432",
		"-U", "finance",
		"-d", "finance_db",
		"-Fc",
		"-f", backupPath,
	)

	cmd.Env = append(
		os.Environ(),
		"PGPASSWORD=masterkey",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("Erro ao executar backup:", err)

		return "", fmt.Errorf(
			"erro ao gerar backup: %w: %s",
			err,
			string(output),
		)
	}

	return backupPath, nil
}

func HandleDBCli(rootPath string, upDB, downDB bool) error {
	ctx := context.Background()

	conn, err := connDB()
	if err != nil {
		logger.Error("Erro ao se conectar ao banco de dados:", err)
		return err
	}

	if err := readContentToRun(rootPath, upDB, downDB); err != nil {
		return err
	}

	if upDB && downDB {
		backupPath, err := backupDB(context.Background(), rootPath)
		if err != nil {
			return err
		}
		logger.Info("Backup realizado em:", backupPath)

		if err := reconstructDB(ctx, conn); err != nil {
			logger.Error("Erro ao reconstruir o banco de dados:", err)
			return err
		}

		return nil
	}

	if upDB && !downDB {
		if err := upDBFunc(ctx, conn); err != nil {
			logger.Error("Erro ao subir o banco de dados:", err)
			return err
		}

		return nil
	}

	if downDB && !upDB {
		backupPath, err := backupDB(context.Background(), rootPath)
		if err != nil {
			return err
		}
		logger.Info("Backup realizado em:", backupPath)

		if err := downDBFunc(ctx, conn); err != nil {
			logger.Error("Erro ao deletar o banco de dados:", err)
			return err
		}

		return nil
	}

	return nil
}

func reconstructDB(ctx context.Context, db *pgxpool.Pool) error {
	if err := downDBFunc(ctx, db); err != nil {
		return err
	}

	if err := upDBFunc(ctx, db); err != nil {
		return err
	}

	logger.Info("Banco de dados reconstituído com sucesso.")
	return nil
}

func downDBFunc(ctx context.Context, db *pgxpool.Pool) error {
	if _, err := db.Exec(
		ctx,
		downContentToRun,
	); err != nil {
		return err
	}

	logger.Info("Banco de dados deletado com sucesso.")
	return nil
}

func upDBFunc(ctx context.Context, db *pgxpool.Pool) error {
	if _, err := db.Exec(
		ctx,
		upContentToRun,
	); err != nil {
		return err
	}

	logger.Info("Banco de dados criado com sucesso.")
	return nil
}
