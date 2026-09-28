package main

import (
	"finance/cmd/cli/routes"
	"finance/internal/config"
	"finance/internal/logger"
	"flag"
	"os"
)

var cfg *config.Config
var wd string

func init() {
	var err error
	wd, err = os.Getwd()

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
	var (
		routeList = flag.Bool("routes", false, "Comando para listar todas as rotas dos módulos")
	)

	flag.Parse()

	if *routeList {
		routes.ListRoutes(wd)
	}
}

/*func main() {


	var (
		path        = flag.String("path", "", "Comando para informar o caminho principal das operações")
		migrateFlag = flag.String("migrate", "", "Comando para rodar a migration. Opções: up, down, fresh, status")
		//seederFlag  = flag.String("seed", "", "Comando para rodar o seeder")

		modelFlag      = flag.String("m", "", "Comando para criação de arquivo padrão da model")
		uuidUse        = flag.Bool("uuid", false, "Comando para criação do model com uuid")
		idUse          = flag.Bool("id", false, "Comando para criação do model com id (int)")
		requests       = flag.Bool("R", false, "Comando para criação da request")
		resource       = flag.Bool("r", false, "Comando para criação de resources")
		seed           = flag.Bool("s", false, "Comando para criação do seeder")
		migrationFlag  = flag.Bool("M", false, "Comando para criação da migration")
		repo           = flag.Bool("repo", false, "Comando para criação de repository pattern")
		controller     = flag.Bool("c", false, "Comando para criação do controller")
		createRepoPath = flag.Bool("create-path", false, "Comando para que caso o caminho do repository pattern não exista, seja criado")
		all            = flag.Bool("a", false, "Comando para separação de pastas por model")
	)

	flag.Parse()

	if *modelFlag != "" {
		commands := scaffold.NewCommandMap(scaffold.CommandOptions{
			UUIDUse:    *uuidUse,
			IDUse:      *idUse,
			Requests:   *requests,
			Resource:   *resource,
			Seed:       *seed,
			Migration:  *migrationFlag,
			Repository: *repo,
			Controller: *controller,
			All:        *all,
		})

		commands["create_repo_path"] = *createRepoPath

		if *repo {
			commands["routes"] = true
			commands["controller"] = true
			commands["service"] = true
		}

		options := scaffold.Options{
			Name:    *modelFlag,
			Command: commands,
			RootDir: *path,
		}

		if err := scaffold.Run(options); err != nil {
			logger.Error("erro ao executar o scaffold", err)
		}

		logger.Info("Scaffold executado com sucesso!")
		return
	}

	if *migrateFlag != "" {
		dbConn, err := database.ConnectDb(cfg.DatabaseConfig)

		if err != nil {
			logger.Error("Erro ao conectar ao banco de dados:", err)
		}

		defer dbConn.Close()

		logger.Info("Banco de dados conectado com sucesso!")

		if err := migration.Run(context.Background(), dbConn, migration.Options{
			Dir:     "database/migration",
			Command: migration.Command(*migrateFlag),
		}); err != nil {
			logger.Error("Erro ao rodar a migration:", err)
		}

		logger.Info("Comando executado com sucesso!")
		return

	}
}

*/
