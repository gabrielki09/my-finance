package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func Init() (*pgxpool.Pool, error) {
	if Pool != nil {
		return Pool, nil
	}

	_ = godotenv.Load()

	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)

	config, err := pgxpool.ParseConfig(dsn)

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
