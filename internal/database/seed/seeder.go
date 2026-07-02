package seed

import (
	"embed"
	"finance/internal/logger"
	categorymodel "finance/models/category"
	financialmodel "finance/models/financial"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"gopkg.in/yaml.v3"
)

//go:embed seeds.yaml
var seedFiles embed.FS

type AccountSeed struct {
	Name string                              `yaml:"name"`
	Type financialmodel.FinancialAccountType `yaml:"type"`
}

type CategorySeed struct {
	Name string                      `yaml:"name"`
	Type categorymodel.CategoryTpyes `yaml:"type"`
}

type SeederFile struct {
	Categories []CategorySeed `yaml:"categories"`
	Accounts   []AccountSeed  `yaml:"categories"`
}

func splitSeeds(s string) []string {
	var splited []string

	splite := strings.Split(s, "-")

	for _, s := range splite {
		splited = append(splited, s)
	}

	return splited
}

func getFileName(s string) string {
	var name string

	switch s {
	case "category":
		name = "category.yaml"
	case "teste":
		name = "teste.yaml"
	}

	return name
}

func parseYAMLFile(seed string) (SeederFile, error) {
	logger.General.Info.Println("Called parseYAMLFile")
	var data SeederFile

	file, err := seedFiles.ReadFile(seed)

	if err != nil {
		return data, err
	}

	if err := yaml.Unmarshal(file, &data); err != nil {
		logger.General.Error.Println("Erro ao associaro o arquivo yaml:", err)
		return data, err
	}

	return data, nil
}

func HandleSeeds(db *pgxpool.Pool, seed *string) (err error) {
	offPointerSeed := strings.TrimSpace(*seed)

	if offPointerSeed == "all" {
		// block for all seeders

	}

	splitedSeeds := splitSeeds(offPointerSeed)

	var collectedSeeds []any

	for _, seed := range splitedSeeds {
		switch seed {
		case "category":
			_, err := parseYAMLFile(getFileName(seed))

			if err != nil {
				return err
			}

		}
	}

	logger.General.Info.Println("Script:", collectedSeeds)
	return err
}
