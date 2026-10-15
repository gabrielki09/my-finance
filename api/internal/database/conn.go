package database

import (
	"context"
	"finance/internal/config"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func ConnectDb(databaseConfig *config.DatabaseConfig) (*pgxpool.Pool, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro carregar a .env:", err)
	}

	if Pool != nil {
		return Pool, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		databaseConfig.DbHost,
		databaseConfig.DbPort,
		databaseConfig.DbUserName,
		databaseConfig.DbPassword,
		databaseConfig.DbDatabase,
	)

	config, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		log.Fatal("Erro ao analisar as config:", err)
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		log.Fatal("Erro validar a conexão com o banco de dados:", err)
		return nil, err
	}

	Pool = pool
	return Pool, nil
}
