package financialobligationrepository

import (
	"context"
	financialobligationrequest "finance/internal/http/request/financial/financial_obligation"
	"finance/internal/logger"
	financialmodel "finance/models/financial"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FinancialObligationRepository struct {
	db *pgxpool.Pool
}

func NewFinancialObligationRepository(db *pgxpool.Pool) *FinancialObligationRepository {
	return &FinancialObligationRepository{
		db: db,
	}
}

func (f FinancialObligationRepository) GetAll(ctx context.Context) ([]financialmodel.FinancialObligationModel, error) {
	var financialObligations []financialmodel.FinancialObligationModel

	financialObligationRows, err := f.db.Query(
		ctx,
		`
			SELECT
				id,
				category_id,
				description,
				type,
				status,
				created_at,
				updated_at
			FROM
				financial_obligations
			WHERE
				deleted_at IS NULL
		`,
	)

	if err != nil {
		logger.General.Error.Println("Erro ao executar o select:", err)
		return []financialmodel.FinancialObligationModel{}, err
	}

	defer financialObligationRows.Close()

	for financialObligationRows.Next() {
		var financialObligation financialmodel.FinancialObligationModel

		if err := financialObligationRows.Scan(
			&financialObligation.Id,
			&financialObligation.CategoryId,
			&financialObligation.Description,
			&financialObligation.Type,
			&financialObligation.Status,
			&financialObligation.CreatedAt,
			&financialObligation.UpdatedAt,
			&financialObligation.DeletedAt,
		); err != nil {
			logger.General.Error.Println("Erro ao ler os dados da consulta:", err)
			return []financialmodel.FinancialObligationModel{}, err
		}

		financialObligations = append(financialObligations, financialObligation)
	}

	if err := financialObligationRows.Err(); err != nil {
		logger.General.Error.Println("Erro ao ler os dados da consulta:", err)
		return []financialmodel.FinancialObligationModel{}, err
	}

	return financialObligations, nil
}

func (f *FinancialObligationRepository) Create(ctx context.Context, payload financialobligationrequest.FinancialObligationRequest) (financialmodel.FinancialObligationModel, error) {
	var financialObligation financialmodel.FinancialObligationModel

	return financialObligation, nil
}

func (f *FinancialObligationRepository) ValidCategoryType(ctx context.Context, categoryId int, movementType financialmodel.FinancialObligationsTpyes) (bool, error) {
	var checkedCategoryType bool

	if err := f.db.QueryRow(
		ctx,
		`
			SELECT
				CASE
					WHEN $2 = 'entry' AND c."type" IN ('income', 'both') THEN TRUE
					WHEN $2 = 'exit' AND c."type" IN ('expense', 'both') THEN TRUE
					ELSE FALSE
				END AS checked_category_type
			FROM

			WHERE
				c.id = $1 AND
				c.active = TRUE AND
				c.deleted_at IS NULL 
		`,
		categoryId,
		movementType,
	).Scan(&checkedCategoryType); err != nil {
		logger.General.Error.Println("Erro ao conferir a categoria")
		return false, err
	}

	return checkedCategoryType, nil
}
