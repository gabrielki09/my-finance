package main

import (
	"context"
	"errors"
	"finance/internal/database"
	"finance/internal/database/migration"
	"finance/internal/database/seed"
	"finance/internal/logger"
	"finance/internal/routes"
	"flag"
	"os"
)

func main() {
	logger.Init()
	ctx := context.Background()

	wd, err := os.Getwd()

	if err != nil {
		logger.General.Error.Fatal("Erro ao obter diretório atual:", err)
	} else {
		logger.General.Info.Println("Rodando em:", wd)
	}

	dbConn, err := database.Init()

	if err != nil {
		logger.General.Error.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	defer dbConn.Close()

	migrateFlag := flag.String("migrate", "", "Rodar migrate")
	seederFlag := flag.String("seed", "", "Rodar seeder")

	flag.Parse()

	if *migrateFlag != "" {
		if err := migration.HandleMigrate(ctx, dbConn, *migrateFlag); err != nil {
			logger.General.Error.Fatal("Erro ao rodar a migration:", err)
		}

		logger.General.Info.Println("Migrate rodada com sucesso!")
		return

	} else if *seederFlag != "" {
		if err := seed.HandleSeeds(ctx, dbConn, seederFlag); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				logger.General.Error.Fatal("Seeder não localizado")
			}

			logger.General.Error.Fatal("Erro ao rodar o seeder:", err)
		}

		logger.General.Info.Println("Seeder rodado com sucesso!")
		return

	}

	logger.General.Info.Println("Banco de dados conectado com sucesso!")
	routes.StartServer(dbConn)
}
