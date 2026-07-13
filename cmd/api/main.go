package main

import (
	"context"
	"finance/internal/database"
	"finance/internal/logger"
	"finance/internal/routes"
	"flag"
	"log"
	"os"

	migration "github.com/gabrielki09/go-scaffold-gk/pkg/migration"
	scaffold "github.com/gabrielki09/go-scaffold-gk/pkg/scaffold"
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

	logger.General.Info.Println("Banco de dados conectado com sucesso!")

	var (
		migrateFlag = flag.String("migrate", "", "Comando para rodar a migration. Opções: up, down, fresh, status")
		seederFlag  = flag.String("seed", "", "Comando para rodar o seeder")

		modelFlag        = flag.String("model", "", "Comando para criação de arquivo padrão da model")
		uuidUse          = flag.Bool("uuid", false, "Comando para criação do model com uuid")
		idUse            = flag.Bool("id", false, "Comando para criação do model com id (int)")
		separateByFolder = flag.Bool("S", false, "Comando para separação de pastas por model")
		requests         = flag.Bool("R", false, "Comando para criação da request")
		resource         = flag.Bool("r", false, "Comando para criação de resources")
		seed             = flag.Bool("s", false, "Comando para criação do seeder")
		migrationFlag    = flag.Bool("m", false, "Comando para criação da migration")
		controller       = flag.Bool("c", false, "Comando para criação do controller")
		all              = flag.Bool("a", false, "Comando para separação de pastas por model")
	)

	flag.Parse()

	flags := make(map[string]bool)

	if *modelFlag != "" {
		if *uuidUse && *idUse {
			log.Fatal("somente um tipo de ID pode ser utilizado.")
		}

		if !*uuidUse && !*idUse {
			log.Fatal("informe o tipo de ID: -uuid ou -id")
		}
	}

	flags["model"] = true
	flags["uuid_use"] = *uuidUse
	flags["id_use"] = *idUse
	flags["separate_by_folder"] = *separateByFolder
	flags["requests"] = *requests
	flags["resource"] = *resource
	flags["seed"] = *seed
	flags["migration"] = *migrationFlag
	flags["controller"] = *controller

	if *all {
		for key := range flags {
			if key == "uuid_use" || key == "id_use" || key == "separate_by_folder" {
				continue
			}

			flags[key] = true
		}
	}

	if *modelFlag != "" {
		options := scaffold.Options{
			Name:             *modelFlag,
			SeparateByFolder: flags["separate_by_folder"],
			Command:          flags,
		}

		if err := scaffold.Run(options); err != nil {
			log.Fatal(err)
		}

		logger.General.Info.Println("Comando executado com sucesso!")
		return
	}

	if *migrateFlag != "" {
		if err := migration.Run(ctx, dbConn, migration.Options{
			Dir:     "database/migration",
			Command: migration.Command(*migrateFlag),
		}); err != nil {
			logger.General.Error.Fatal("Erro ao rodar a migration:", err)
		}

		logger.General.Info.Println("Comando executado com sucesso!")
		return

	} else if *seederFlag != "" {

	}

	routes.StartServer(dbConn)
}
