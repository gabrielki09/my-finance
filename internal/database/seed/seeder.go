package seed

import (
	"context"
	"embed"
	"finance/internal/logger"
	categorymodel "finance/models/category"
	financialmodel "finance/models/financial"
	"strings"

	"github.com/jackc/pgx/v5"
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
	Name  string                      `yaml:"name"`
	Type  categorymodel.CategoryTpyes `yaml:"type"`
	Color string                      `yaml:"color"`
}

type SeederFile struct {
	Accounts   []AccountSeed  `yaml:"accounts"`
	Categories []CategorySeed `yaml:"categories"`
}

func splitSeeds(s string) (splited []string) {
	splite := strings.Split(s, "-")

	for _, s := range splite {
		splited = append(splited, s)
	}

	return splited
}

func parseYAMLFile() (SeederFile, error) {
	logger.General.Info.Println("Called parseYAMLFile")
	var data SeederFile

	file, err := seedFiles.ReadFile("seeds.yaml")

	if err != nil {
		return data, err
	}

	if err := yaml.Unmarshal(file, &data); err != nil {
		logger.General.Error.Println("Erro ao associaro o arquivo yaml:", err)
		return data, err
	}

	return data, nil
}

func HandleSeeds(ctx context.Context, db *pgxpool.Pool, seed *string) (err error) {
	if seed == nil {
		return nil
	}

	offPointerSeed := strings.TrimSpace(*seed)

	data, err := parseYAMLFile()
	if err != nil {
		return err
	}

	var splitedSeeds []string

	tx, err := db.Begin(ctx)

	if err != nil {
		logger.General.Error.Println("Erro ao iniciar a transação:", err)
		return err
	}

	if offPointerSeed == "all" {
		if err := runAll(ctx, data, tx); err != nil {
			return err
		}

		if err := tx.Commit(ctx); err != nil {
			logger.General.Error.Println("Erro ao commitar a transação:", err)
			return err
		}

		return err

	} else {
		splitedSeeds = splitSeeds(offPointerSeed)
	}

	for _, seed := range splitedSeeds {
		switch seed {
		case "account":
			if err := runAccount(ctx, data.Accounts, tx); err != nil {
				return err
			}

			if err := tx.Commit(ctx); err != nil {
				logger.General.Error.Println("Erro ao commitar a transação:", err)
				return err
			}
		case "category":
			if err := runCategory(ctx, data.Categories, tx); err != nil {
				return err
			}

			if err := tx.Commit(ctx); err != nil {
				logger.General.Error.Println("Erro ao commitar a transação:", err)
				return err
			}
		}
	}

	return err
}

func runAll(ctx context.Context, items SeederFile, tx pgx.Tx) (err error) {
	if err = runCategory(ctx, items.Categories, tx); err != nil {
		return err
	}

	if err = runAccount(ctx, items.Accounts, tx); err != nil {
		return err
	}

	return err
}

func runAccount(ctx context.Context, items []AccountSeed, tx pgx.Tx) (err error) {
	for _, item := range items {
		if _, err := tx.Exec(
			ctx,
			`
				INSERT INTO financial_accounts (name, type)
				VALUES 	($1, $2)	
			`,
			item.Name,
			item.Type,
		); err != nil {
			logger.General.Error.Println("Erro ao realizar o insert:", err)
			return err
		}
	}

	return err
}

func runCategory(ctx context.Context, items []CategorySeed, tx pgx.Tx) (err error) {
	for _, item := range items {
		if _, err := tx.Exec(
			ctx,
			`
				INSERT INTO categories (name, type, color)
				VALUES 	($1, $2, $3)	
			`,
			item.Name,
			item.Type,
			item.Color,
		); err != nil {
			logger.General.Error.Println("Erro ao realizar o insert:", err)
			return err
		}
	}

	return err
}
