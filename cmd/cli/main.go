package main

import (
	"finance/cmd/cli/db"
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
		routeListFlag = flag.Bool("routes", false, "Comando para listar todas as rotas dos módulos.")

		dbFlag    = flag.Bool("db", false, "Comando para operações do banco.")
		dbUpFlag  = flag.Bool("up", false, "Comando usado com a flag db para rodar o create do banco.")
		dbDowFlag = flag.Bool("down", false, "Comando usado com a flag db para derrubar todo o create do banco.")
	)

	flag.Parse()

	if *routeListFlag {
		if err := routes.ListRoutes(wd); err != nil {
			logger.Fatal("Erro ao executar route cli: ", err)
		}
	}

	if *dbFlag {
		if err := db.HandleDBCli(wd, *dbUpFlag, *dbDowFlag); err != nil {
			logger.Fatal("Erro ao executar db cli: ", err)
		}
	}
}
