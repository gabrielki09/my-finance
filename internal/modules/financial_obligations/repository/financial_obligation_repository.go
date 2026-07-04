package financialobligationrepository

import (
	"context"
	"errors"
	"finance/internal/apperrors"
	transactionhelper "finance/internal/helpers/transaction"
	"finance/internal/helpers/utils"
	financialobligationrequest "finance/internal/http/request/financial/financial_obligation"
	"finance/internal/logger"
	financialmodel "finance/models/financial"
	"fmt"

	"github.com/jackc/pgx/v5"
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

		if errors.Is(err, pgx.ErrNoRows) {
			return financialObligation, apperrors.ErrNotFound
		}

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

func (f *FinancialObligationRepository) genericInsertObligationSettlements(
	ctx context.Context,
	tx pgx.Tx,
	payload financialobligationrequest.PayFinancialObligationRequest,
	financialTransactionId int,
) (err error) {
	if _, err := tx.Exec(
		ctx,
		`
			INSERT INTO obligation_settlements
				(
					obligation_id,
					transaction_id,
					amount
				)
			VALUES
				(
					$1,
					$2, 
					$3
				)
		`,
		payload.FinancialObligationId,
		financialTransactionId,
		payload.AmountPaid,
	); err != nil {
		logger.General.Error.Println("Erro ao criar o registro de pagamento da obrigação financeira:", err)
		return err
	}

	return err
}

func (f *FinancialObligationRepository) genericInsertFinancialTransactions(
	ctx context.Context,
	tx pgx.Tx,
	payload financialobligationrequest.PayFinancialObligationRequest,
	financialObligation financialmodel.FinancialObligationModel,
) (id int, err error) {
	outstandingBalance, err := f.CalculateOutstandingBalance(ctx, financialObligation.Id)

	if err != nil {
		return id, err
	}

	description := utils.Ternary(
		payload.AmountPaid == outstandingBalance,
		fmt.Sprintf("Pagamento parcial da obrigação finaceira N° %d ", financialObligation.Id),
		fmt.Sprintf("Pagamento da obrigação finaceira N° %d ", financialObligation.Id),
	)

	logger.General.Info.Println("description:", description)

	if err := tx.QueryRow(
		ctx,
		`
			INSERT INTO financial_transactions
				(
					financial_account_id,
					category_id,
					reversed_transaction_id,
					description,
					movement_type,
					operation_type,
					amount,
					movement_date,
					reference_date,
					origin_type,
					origin_id
				)
			VALUES
				(
					$1,
					$2,
					$3,
					$4,
					$5,
					$6,
					$7,
					$8,
					$9,
					$10,
					$11
				)
			RETURNING
				id
		`,
		payload.FinancialAccountId,
		financialObligation.CategoryId,
		nil,
		description,
		financialmodel.EXIT,
		financialmodel.ORIGINAL,
		payload.AmountPaid,
		payload.PaymentDate,
		financialObligation.DueDate,
		"financial_obligation",
		payload.FinancialObligationId,
	).Scan(&id); err != nil {
		logger.General.Error.Println("Erro ao criar o registro da transação do pagamento da obrigação financeira:", err)
		return id, err
	}

	return id, err
}

func (f *FinancialObligationRepository) selectFinancialObligationForUpdate(
	ctx context.Context,
	tx pgx.Tx,
	financialObligationId int,
) (financialObligation financialmodel.FinancialObligationModel, err error) {
	if err := tx.QueryRow(
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
				notes
			FROM	
				financial_obligations
			WHERE
				id = $1	AND
				deleted_at IS NULL
			FOR UPDATE
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
	); err != nil {
		logger.General.Error.Println("Erro ao localizar a obrigação financeiro pelo ID:", err)
		return financialObligation, err
	}
	return financialObligation, err
}

func (f *FinancialObligationRepository) genericUpdateFinancialObligation(
	ctx context.Context,
	tx pgx.Tx,
	amountPaid float64,
	financialObligationId int,
) (err error) {
	outstandingBalance, err := f.CalculateOutstandingBalance(ctx, financialObligationId)

	if err != nil {
		return err
	}

	if _, err := tx.Exec(
		ctx,
		`
			UPDATE 
				financial_obligations
			SET
				status = $2,
				updated_at = now()
			WHERE
				id = $1
		`,
		financialObligationId,
		utils.Ternary(
			amountPaid == outstandingBalance,
			financialmodel.SETTLED,
			financialmodel.PARTIALLY_SETTLED,
		),
	); err != nil {
		logger.General.Error.Println("Erro ao definir o status como pago:", err)
		return err
	}

	return err
}

func (f *FinancialObligationRepository) CalculateOutstandingBalance(
	ctx context.Context,
	financialObligationId int,
) (amount float64, err error) {
	if err := f.db.QueryRow(
		ctx,
		`
			SELECT 
				fo.original_amount 
				-
				COALESCE(SUM(os.amount), 0) AS outstanding_balance
			FROM
				financial_obligations fo 
			LEFT JOIN obligation_settlements os 
				ON os.obligation_id = fo.id
				AND os.canceled_at IS NULL
			WHERE
				fo.id = $1 
				AND fo.deleted_at IS null
			GROUP BY
				fo.id,
				fo.original_amount
		`,
		financialObligationId,
	).Scan(&amount); err != nil {
		logger.General.Error.Println("Erro ao conferir o total pendente: ", err)

		if errors.Is(err, pgx.ErrNoRows) {
			return amount, err
		}

		return amount, err
	}

	return amount, err
}

func (f *FinancialObligationRepository) Pay(
	ctx context.Context,
	payload financialobligationrequest.PayFinancialObligationRequest,
) error {
	if err := transactionhelper.WithTransaction(
		ctx,
		f.db,
		func(tx pgx.Tx) error {
			financialObligation, err := f.selectFinancialObligationForUpdate(ctx, tx, payload.FinancialObligationId)

			if err != nil {
				logger.General.Error.Println("Erro ao localizar a obrigação financeira que será paga:", err)
				return err
			}

			logger.General.Info.Println("Obrigação financeira que será paga:", financialObligation)

			logger.General.Info.Println("O valor pago é igual o da obrigação financeira:", payload.AmountPaid)

			financialTransactionId, err := f.genericInsertFinancialTransactions(
				ctx,
				tx,
				payload,
				financialObligation,
			)

			if err != nil {
				logger.General.Error.Println("Erro ao cadastrar a transação financeira: ", err)
				return err
			}

			if err := f.genericInsertObligationSettlements(
				ctx,
				tx,
				payload,
				financialTransactionId,
			); err != nil {
				logger.General.Error.Println("Erro ao cadastrar a liquidição financeira: ", err)
				return err
			}

			if err := f.genericUpdateFinancialObligation(
				ctx,
				tx,
				payload.AmountPaid,
				payload.FinancialObligationId,
			); err != nil {
				logger.General.Error.Println("Erro ao alterar o status da obrigação financeira: ", err)
				return err
			}

			return nil
		}); err != nil {

		logger.General.Error.Println("Erro durante a transação de pagamento:", err)
		return err

	}

	return nil

}
