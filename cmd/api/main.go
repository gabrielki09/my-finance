package main

import (
	"context"
	"errors"
	"finance/internal/constants/colors"
	"finance/internal/database"
	"finance/internal/database/migration"
	"finance/internal/database/seed"
	"finance/internal/logger"
	"finance/internal/routes"
	"flag"
	"log"
	"os"
)

func main() {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal("Erro ao obter diretório atual:", err)
	} else {
		log.Println("Rodando em:", wd)
	}

	ctx := context.Background()
	logger.Init()

	dbConn, err := database.Init()

	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	defer dbConn.Close()

	migrateFlag := flag.Bool("migrate", false, "Rodar migrate")
	seederFlag := flag.String("seed", "", "Rodar seeder")

	flag.Parse()

	if *migrateFlag {
		if err := migration.HandleMigrate(ctx, dbConn); err != nil {
			logger.General.Error.Fatal(colors.Red+"Erro ao rodar a migration:", err)
		}

		logger.General.Info.Println(colors.Green + "Migrate rodada com sucesso!")
		return

	} else if *seederFlag != "" {
		if err := seed.HandleSeeds(ctx, dbConn, seederFlag); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				logger.General.Error.Fatal(colors.Red + "Seeder não localizado")
			}

			logger.General.Error.Fatal(colors.Red+"Erro ao rodar o seeder:", err)
		}

		logger.General.Info.Println(colors.Green + "Seeder rodado com sucesso!")
		return

	}

	logger.General.Info.Println(colors.Green + "Banco de dados conectado com sucesso!")
	routes.StartServer(dbConn)
}
