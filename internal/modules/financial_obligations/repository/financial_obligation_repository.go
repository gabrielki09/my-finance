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

	logger.General.Info.Println("Query:", query)

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
	logger.General.Info.Println("FinancialObligationRepository - Update")

	financialObligation, err := f.FindById(ctx, financialObligationId)

	if err != nil {
		logger.General.Error.Println("Erro ao localizar a obrigação financeira:", err)
		return financialmodel.FinancialObligationModel{}, err
	}

	logger.General.Info.Println("financialObligation:", financialObligation)

	switch financialObligation.Status {
	case financialmodel.PARTIALLY_SETTLED:
		logger.General.Info.Println("Status atual: ", financialmodel.PARTIALLY_SETTLED)

		if err := f.db.QueryRow(
			ctx,
			`
				UPDATE 	
					financial_obligations
				SET
					description = $2,
					notes = $3
				WHERE 
					id = $1 AND
					status = 'partially_settled'
				RETURNING
					id,
					category_id,
					description,
					type,
					status,
					original_amount,
					due_date,
					competence_date,
					notes
			`,
			financialObligationId,
			payload.Description,
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
		); err != nil {
			return financialmodel.FinancialObligationModel{}, err
		}

	case financialmodel.PENDING:
		logger.General.Info.Println("Status atual: ", financialmodel.PENDING)

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
				RETURNING
					id,
					category_id,
					description,
					type,
					status,
					original_amount,
					due_date,
					competence_date,
					notes
			`,
			financialObligationId,
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
		); err != nil {
			return financialmodel.FinancialObligationModel{}, err
		}
	}

	return financialObligation, nil
}

func (f *FinancialObligationRepository) ValidCategoryType(ctx context.Context, categoryId int, obligationType financialmodel.FinancialObligationsTypes) (bool, error) {
	logger.General.Info.Println("FinancialObligationRepository - ValidCategoryType")
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
	logger.General.Info.Println("FinancialObligationRepository - FindById")

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
				financial_obligations
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
		logger.General.Error.Println("Erro ao localizar a obrigação financeiro pelo ID:", err)
		return financialObligation, err
	}

	return financialObligation, nil

}

func (f *FinancialObligationRepository) Cancel(ctx context.Context, financialObligationId int) error {
	if _, err := f.db.Exec(
		ctx,
		`
			UPDATE 
				financial_obligations
			SET
				status = 'canceled',
				canceled_at = now()
			WHERE
				id = $1
		`,
		financialObligationId,
	); err != nil {
		return err
	}

	return nil

}
