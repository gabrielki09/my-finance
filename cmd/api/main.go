package main

import (
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
		if err := migration.RunMigrate(dbConn); err != nil {
			log.Fatal(colors.Red+"Erro ao rodar a migration:", err)
		}

		log.Println(colors.Green + "Migrate rodada com sucesso!")
		return

	} else if *seederFlag != "" {
		if err := seed.HandleSeeds(dbConn, seederFlag); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				log.Fatal(colors.Red + "Seeder não localizado")
			}

			log.Fatal(colors.Red+"Erro ao rodar o seeder:", err)
		}

		log.Println(colors.Green + "Seeder rodado com sucesso!")
		return

	}

	log.Println("Banco de dados conectado com sucesso!")
	routes.StartServer(dbConn)
}
