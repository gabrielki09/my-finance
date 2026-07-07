package main

import (
	"context"
	"errors"
	"finance/internal/database"
	"finance/internal/database/seed"
	"finance/internal/logger"
	"finance/internal/routes"
	"flag"
	"os"

	"github.com/GabrielK09/go-migrate-gk/migration"
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

	migrateFlag := flag.String("migrate", "", "Rodar migration. Opções: up, down, fresh, status")
	seederFlag := flag.String("seed", "", "Rodar seeder")

	flag.Parse()

	extraArgs := flag.Args()

	if *migrateFlag != "" {

		if *migrateFlag == "create" {
			if len(extraArgs) == 0 {
				logger.General.Error.Fatal("Informe o nome da migration. Exemplo: -migrate=create create_users")
			}

			migrationName := extraArgs[0]
		}

		if err := migration.Run(ctx, dbConn, migration.Options{
			Dir:     "database/migration",
			Command: migration.Command(*migrateFlag),
		}); err != nil {
			logger.General.Error.Fatal("Erro ao rodar a migration:", err)
		}

		logger.General.Info.Println("Migrate executada com sucesso!")
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
