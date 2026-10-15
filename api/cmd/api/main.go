package main

import (
	"finance/internal/config"
	"finance/internal/database"
	"finance/internal/logger"
	"finance/internal/routes"
	"os"
)

var cfg *config.Config

func init() {
	wd, err := os.Getwd()

	if err != nil {
		logger.Error("Erro ao obter diretório atual:", err)
	} else {
		logger.Info("Rodando em:", "local", wd)
	}

	cfg, err = config.LoadConfig()
	if err != nil {
		logger.Error("Erro ao conectar ao banco de dados:", err)
	}
}

func main() {
	dbConn, err := database.ConnectDb(cfg.DatabaseConfig)
	if err != nil {
		logger.Error("Erro ao conectar ao banco de dados:", err)
	}

	defer dbConn.Close()

	logger.Info("Banco de dados conectado com sucesso!")

	routes.StartServer(dbConn, cfg.ApiConfig)
}
