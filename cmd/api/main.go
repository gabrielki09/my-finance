package main

import (
	"finance/internal/database"
	"finance/internal/logger"
	"finance/internal/routes"
	"log"
	"os"
)

func main() {
	logger.Init()

	wd, err := os.Getwd()

	if err != nil {
		log.Fatal("Erro ao obter diretório atual:", err)
	} else {
		log.Println("Rodando em:", wd)
	}

	dbConn, err := database.Init()

	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	defer dbConn.Close()

	log.Println("Banco de dados conectado com sucesso!")
	routes.StartServer(dbConn)
}
