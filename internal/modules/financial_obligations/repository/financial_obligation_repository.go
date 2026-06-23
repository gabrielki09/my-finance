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

func (f FinancialObligationRepository) GetAll(ctx context.Context, query string, args []any) ([]financialmodel.FinancialObligationModel, error) {
	var financialObligations []financialmodel.FinancialObligationModel

	financialObligationRows, err := f.db.Query(
		ctx,
		query,
		args...,
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
			&financialObligation.OriginalAmount,
			&financialObligation.DueDate,
			&financialObligation.CompetenceDate,
			&financialObligation.Notes,
			&financialObligation.CreatedAt,
			&financialObligation.UpdatedAt,
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

	if err := f.db.QueryRow(
		ctx,
		`
			INSERT INTO financial_obligations
				(
					category_id,
					description,
					type,
					original_amount,
					due_date,
					competence_date,
					notes
				)

			VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7
			)

			RETURNING
				id,
				category_id,
				description,
				type,
				status,
				original_amount,
				due_date,
				competence_date,
				notes,
				created_at,
				updated_at
		`,
		payload.CategoryId,
		payload.Description,
		payload.Type,
		payload.OriginalAmount,
		payload.DueDate,
		payload.CompetenceDate,
		payload.Notes,
	).Scan(
		&financialObligation.Id,
		&financialObligation.CategoryId,
		&financialObligation.Description,
		&financialObligation.Type,
		&financialObligation.Status,
		&financialObligation.OriginalAmount,
		&financialObligation.DueDate,
		&financialObligation.CompetenceDate,
		&financialObligation.Notes,
		&financialObligation.CreatedAt,
		&financialObligation.UpdatedAt,
	); err != nil {
		return financialmodel.FinancialObligationModel{}, err
	}

	return financialObligation, nil
}

func (f *FinancialObligationRepository) Update(ctx context.Context, payload financialobligationrequest.FinancialObligationRequest, financialObligationId int) (financialmodel.FinancialObligationModel, error) {
	financialObligation, err := f.FindById(ctx, financialObligationId)

	if err != nil {
		return financialmodel.FinancialObligationModel{}, nil
	}

	switch financialObligation.Status {
	case financialmodel.PARTIALLY_SETTLED:
		if err := f.db.QueryRow(
			ctx,
			`
				UPDATE 	
					financial_obligations
				SET
					description = $2,
					notes = $3
				WHERE
					id = $1
					status = 'partially_settled'
			`,
			financialObligationId,
			payload.Description,
			payload.Notes,
		).Scan(); err != nil {
			return financialmodel.FinancialObligationModel{}, nil
		}

	case financialmodel.PENDING:
		if err := f.db.QueryRow(
			ctx,
			`
				UPDATE 	
					financial_obligations
				SET
					category_id = $2,
					description = $3,
					type = $4,
					original_amount = $5,
					due_date = $6,
					competence_date = $7,
					notes = $8
				WHERE
					id = $1 AND
					status = 'pending'
			`,
			financialObligationId,
			payload.CategoryId,
			payload.Description,
			payload.Type,
			payload.OriginalAmount,
			payload.DueDate,
			payload.CompetenceDate,
			payload.Notes,
		).Scan(); err != nil {
			return financialmodel.FinancialObligationModel{}, nil
		}
	}

	return financialObligation, nil
}

func (f *FinancialObligationRepository) ValidCategoryType(ctx context.Context, categoryId int, obligationType financialmodel.FinancialObligationsTypes) (bool, error) {
	var checkedCategoryType bool

	if err := f.db.QueryRow(
		ctx,
		`
			SELECT EXISTS (
				SELECT
					1
				FROM
					categories c
				WHERE
					c.id = $1 AND
					c.deleted_at IS NULL
					AND (
						($2 = 'receivable' AND c."type" IN ('income', 'both'))
						OR
						($2 = 'payable' AND c."type" IN ('expense', 'both'))
					)
					
			) AS checked_category_type
		`,
		categoryId,
		obligationType,
	).Scan(&checkedCategoryType); err != nil {
		logger.General.Error.Println("Erro ao conferir a categoria:", err)
		return false, err
	}

	logger.General.Error.Println("Categoria é válida?", checkedCategoryType)

	return checkedCategoryType, nil
}

func (f *FinancialObligationRepository) FindById(ctx context.Context, financialObligationId int) (financialmodel.FinancialObligationModel, error) {
	var financialObligation financialmodel.FinancialObligationModel

	if err := f.db.QueryRow(
		ctx,
		`
			SELECT
				id,
				category_id,
				description,
				type,
				status,
				original_amount,
				due_date,
				competence_date,
				notes,
				created_at,
				updated_at
			FROM

			WHERE
				id = $1
		`,
		financialObligationId,
	).Scan(
		&financialObligation.Id,
		&financialObligation.CategoryId,
		&financialObligation.Description,
		&financialObligation.Type,
		&financialObligation.Status,
		&financialObligation.OriginalAmount,
		&financialObligation.DueDate,
		&financialObligation.CompetenceDate,
		&financialObligation.Notes,
		&financialObligation.CreatedAt,
		&financialObligation.UpdatedAt,
	); err != nil {
		return financialObligation, nil
	}

	return financialObligation, nil

}
