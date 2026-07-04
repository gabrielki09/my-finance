package financialtransactionrepository

import (
	"context"
	"errors"
	"finance/internal/apperrors"
	transactionhelper "finance/internal/helpers/transaction"
	financialtransactionrequest "finance/internal/http/request/financial/financial_transaction"
	"finance/internal/logger"
	financialmodel "finance/models/financial"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FinancialTransactionRepository struct {
	db *pgxpool.Pool
}

func NewFinancialTransactionRepository(db *pgxpool.Pool) *FinancialTransactionRepository {
	return &FinancialTransactionRepository{
		db: db,
	}
}

func (f *FinancialTransactionRepository) GetAll(ctx context.Context) ([]financialmodel.FinancialTransactionsModel, error) {
	var financialTransactions []financialmodel.FinancialTransactionsModel

	financialTransactionsRows, err := f.db.Query(
		ctx,
		`
			SELECT
				id,
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
				origin_id,
				idempotency_key,
				created_at,
				canceled_at
			FROM
				financial_transactions
			WHERE
				canceled_at IS NULL
		`,
	)

	if err != nil {
		logger.General.Error.Println("Erro ao executar o select:", err)
		return []financialmodel.FinancialTransactionsModel{}, err
	}

	defer financialTransactionsRows.Close()

	for financialTransactionsRows.Next() {
		var financialTransaction financialmodel.FinancialTransactionsModel

		if err := financialTransactionsRows.Scan(
			&financialTransaction.Id,
			&financialTransaction.FinancialAccountId,
			&financialTransaction.CategoryId,
			&financialTransaction.ReversedTransactionId,
			&financialTransaction.Description,
			&financialTransaction.MovementType,
			&financialTransaction.OperationType,
			&financialTransaction.Amount,
			&financialTransaction.MovementDate,
			&financialTransaction.ReferenceDate,
			&financialTransaction.OriginType,
			&financialTransaction.OriginId,
			&financialTransaction.IdempotencyKey,
			&financialTransaction.CreatedAt,
			&financialTransaction.CanceledAt,
		); err != nil {
			logger.General.Error.Println("Erro ao ler os dados do select:", err)
			return []financialmodel.FinancialTransactionsModel{}, err
		}

		financialTransactions = append(financialTransactions, financialTransaction)
	}

	return financialTransactions, nil

}

func (f *FinancialTransactionRepository) FindById(ctx context.Context, financialTransactionId int) (financialmodel.FinancialTransactionsModel, error) {
	var financialTransaction financialmodel.FinancialTransactionsModel

	if err := f.db.QueryRow(
		ctx,
		`
			SELECT
				id,
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
				origin_id,
				idempotency_key,
				created_at,
				canceled_at
			FROM
				financial_transactions
			WHERE
				id = $1
		`,
		financialTransactionId,
	).Scan(
		&financialTransaction.Id,
		&financialTransaction.FinancialAccountId,
		&financialTransaction.CategoryId,
		&financialTransaction.ReversedTransactionId,
		&financialTransaction.Description,
		&financialTransaction.MovementType,
		&financialTransaction.OperationType,
		&financialTransaction.Amount,
		&financialTransaction.MovementDate,
		&financialTransaction.ReferenceDate,
		&financialTransaction.OriginType,
		&financialTransaction.OriginId,
		&financialTransaction.IdempotencyKey,
		&financialTransaction.CreatedAt,
		&financialTransaction.CanceledAt,
	); err != nil {
		logger.General.Error.Println("Erro ao ler os dados do select:", err)

		if errors.Is(err, pgx.ErrNoRows) {
			logger.General.Error.Println("O registro não foi localizado:", err)
			return financialmodel.FinancialTransactionsModel{}, apperrors.ErrNotFound

		}

		return financialmodel.FinancialTransactionsModel{}, err
	}

	return financialTransaction, nil
}

func (f *FinancialTransactionRepository) FindByKey(ctx context.Context, key string) (financialmodel.FinancialTransactionsModel, error) {
	var financialTransaction financialmodel.FinancialTransactionsModel

	if err := f.db.QueryRow(
		ctx,
		`
			SELECT
				id,
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
				origin_id,
				idempotency_key,
				created_at,
				canceled_at
			FROM
				financial_transactions
			WHERE
				idempotency_key = $1
		`,
		key,
	).Scan(
		&financialTransaction.Id,
		&financialTransaction.FinancialAccountId,
		&financialTransaction.CategoryId,
		&financialTransaction.ReversedTransactionId,
		&financialTransaction.Description,
		&financialTransaction.MovementType,
		&financialTransaction.OperationType,
		&financialTransaction.Amount,
		&financialTransaction.MovementDate,
		&financialTransaction.ReferenceDate,
		&financialTransaction.OriginType,
		&financialTransaction.OriginId,
		&financialTransaction.IdempotencyKey,
		&financialTransaction.CreatedAt,
		&financialTransaction.CanceledAt,
	); err != nil {
		logger.General.Error.Println("Erro ao ler os dados do select:", err)

		if errors.Is(err, pgx.ErrNoRows) {
			logger.General.Error.Println("O registro não foi localizado:", err)
			return financialmodel.FinancialTransactionsModel{}, apperrors.ErrNotFound

		}

		return financialmodel.FinancialTransactionsModel{}, err
	}

	return financialTransaction, nil
}

func (f *FinancialTransactionRepository) CreateMovement(ctx context.Context, payload financialtransactionrequest.FinancialTransactionRequest) (financialTransaction financialmodel.FinancialTransactionsModel, err error) {

	err = transactionhelper.WithTransaction(
		ctx,
		f.db,
		func(tx pgx.Tx) error {
			if err := tx.QueryRow(
				ctx,
				`
					INSERT INTO financial_transactions (
						financial_account_id,
						category_id,
						description,
						movement_type,
						operation_type,
						amount,
						movement_date,
						reference_date
					) VALUES(
						$1,
						$2,
						$3,
						$4,
						$5,
						$6,
						$7,
						$8
					) RETURNING 
						id,
						financial_account_id,
						category_id,
						description,
						movement_type,
						operation_type,
						amount,
						movement_date,
						reference_date,
						origin_type,
						origin_id,
						idempotency_key,
						created_at,
						canceled_at
				`,
				payload.FinancialAccountId,
				payload.CategoryId,
				payload.Description,
				payload.MovementType,
				payload.OperationType,
				payload.Amount,
				payload.MovementDate,
				payload.ReferenceDate,
			).Scan(
				&financialTransaction.Id,
				&financialTransaction.FinancialAccountId,
				&financialTransaction.CategoryId,
				&financialTransaction.Description,
				&financialTransaction.MovementType,
				&financialTransaction.OperationType,
				&financialTransaction.Amount,
				&financialTransaction.MovementDate,
				&financialTransaction.ReferenceDate,
				&financialTransaction.OriginType,
				&financialTransaction.OriginId,
				&financialTransaction.IdempotencyKey,
				&financialTransaction.CreatedAt,
				&financialTransaction.CanceledAt,
			); err != nil {
				logger.General.Error.Println("Erro ao fazer o insert da transação financeira:", err)
				return err
			}

			// Quando o movement_type == entry for de entrada, o valora da conta financeira precisará pegar o valor atual + o valor informado
			if financialTransaction.MovementType == "entry" {
				if _, err := tx.Exec(
					ctx,
					`
						UPDATE 
							financial_accounts
						SET
							initial_balance = initial_balance + $2
						WHERE
							id = $1
					`,
					payload.FinancialAccountId,
					payload.Amount,
				); err != nil {
					logger.General.Error.Println("Erro ao alterar o agregar o valor da conta financeira:", err)

					return err
				}
			}

			// Quando o movement_type == exit for de saída, o valora da conta financeira precisará pegar o valor atual - o valor informado
			if financialTransaction.MovementType == "exit" {
				if _, err := tx.Exec(
					ctx,
					`
						UPDATE 
							financial_accounts
						SET
							initial_balance = initial_balance - $2
						WHERE
							id = $1
					`,
					payload.FinancialAccountId,
					payload.Amount,
				); err != nil {
					logger.General.Error.Println("Erro ao alterar o descontar o valor da conta financeira:", err)

					return err
				}
			}

			return nil
		})

	return financialTransaction, err
}

func (f *FinancialTransactionRepository) CancelMovement(ctx context.Context, payload financialtransactionrequest.CancelFinancialTransactionRequest) (err error) {

	isNotCanceled, err := f.CheckIsNotCanceled(ctx, payload)

	if !isNotCanceled {
		return fmt.Errorf("Transação já cancelada.")
	}

	err = transactionhelper.WithTransaction(
		ctx,
		f.db,
		func(tx pgx.Tx) error {
			var financialTransaction financialmodel.FinancialTransactionsModel

			if payload.Id != nil {
				financialTransaction, _ = f.FindById(ctx, *payload.Id)

			} else if (payload.IdempotencyKey) != nil {
				financialTransaction, _ = f.FindByKey(ctx, *payload.IdempotencyKey)
			}

			if financialTransaction.MovementType == "entry" {
				logger.General.General.Println("A operação que está sendo cancelada é uma operação de entrada, vai descontar o valor da conta financeira referenciada.")

				if _, err := tx.Exec(
					ctx,
					`
						UPDATE
							financial_accounts
						SET
							initial_balance = initial_balance - $2
						WHERE
							id = $1
					`,
					financialTransaction.Id,
					financialTransaction.Amount,
				); err != nil {
					logger.General.Error.Println("Erro ao atualizar o saldo da conta financeira:", err)

					return err
				}
			}

			if financialTransaction.MovementType == "exit" {
				logger.General.General.Println("A operação que está sendo cancelada é uma operação de saída, vai acrescentar o valor da conta financeira referenciada.")
				if _, err := tx.Exec(
					ctx,
					`
						UPDATE
							financial_accounts
						SET
							initial_balance = initial_balance + $2
						WHERE
							id = $1
					`,
					financialTransaction.Id,
					financialTransaction.Amount,
				); err != nil {
					logger.General.Error.Println("Erro ao atualizar o saldo da conta financeira:", err)

					return err
				}
			}

			if _, err := tx.Exec(
				ctx,
				`
					INSERT INTO financial_transactions (
						financial_account_id,
						category_id,
						reversed_transaction_id,
						description,
						movement_type,
						operation_type,
						amount,
						movement_date,
						reference_date
					) VALUES(
						$1,
						$2,
						$3,
						$4,
						$5,
						$6,
						$7,
						$8,
						$9
					)
				`,
				financialTransaction.FinancialAccountId,
				financialTransaction.CategoryId,
				financialTransaction.Id,
				financialTransaction.Description,
				financialTransaction.MovementType,
				financialTransaction.OperationType,
				financialTransaction.Amount,
				financialTransaction.MovementDate,
				financialTransaction.ReferenceDate,
			); err != nil {
				logger.General.Error.Println("Erro ao fazer o insert do cancelamento")

				return err
			}

			return nil
		})

	return err
}

func (f *FinancialTransactionRepository) ValidCategoryType(ctx context.Context, categoryId int, movementType financialmodel.FinancialTransactionsMovementType) (bool, error) {
	var checkedCategoryType bool
	var exists bool

	if err := f.db.QueryRow(
		ctx,
		`
			SELECT EXISTS
				(
					SELECT
						id
					FROM
						categories
					WHERE
						id = $1
				)
		`,
		categoryId,
	).Scan(
		&exists,
	); err != nil {

		return false, err
	}

	switch movementType {
	case financialmodel.ENTRY:
		if err := f.db.QueryRow(
			ctx,
			`
				SELECT
					CASE 
						WHEN c."type" = 'expense' THEN true
						ELSE false
					END AS is_expense
				FROM 
					categories c 
				WHERE
					id = $1
			`,
			categoryId,
		).Scan(&checkedCategoryType); err != nil {
			logger.General.Error.Println("Erro ao conferir se a categoria é válida para a operação")
			return false, err
		}
	case financialmodel.EXIT:
		if err := f.db.QueryRow(
			ctx,
			`
				SELECT
					CASE 
						WHEN c."type" = 'income' THEN true
						ELSE false
					END AS is_income
				FROM 
					categories c 
				WHERE
					id = $1
			`,
			categoryId,
		).Scan(&checkedCategoryType); err != nil {
			logger.General.Error.Println("Erro ao conferir se a categoria é válida para a operação")
			return false, err
		}
	default:
		return false, fmt.Errorf("Tipo de movimento financeiro inválido.")
	}

	return checkedCategoryType, nil
}

func (f *FinancialTransactionRepository) ValidateIsSameIdAndIdempotencyKey(ctx context.Context, financialTransactionId int, idempotencyKey string) (exists bool, err error) {
	logger.General.Info.Println("Called ValidateIsSameIdAndIdempotencyKey")

	if err := f.db.QueryRow(
		ctx,
		`
			SELECT EXISTS 
				(
					SELECT
						1
					FROM
						financial_transactions
					WHERE
						(
							id = $1 
							OR idempotency_key = $2
						)
				)
		`,
		financialTransactionId,
		idempotencyKey,
	).Scan(
		&exists,
	); err != nil {
		logger.General.Error.Println("Erro ao conferir se a transação existe:", err)
		return exists, err
	}

	return exists, err
}

func (f *FinancialTransactionRepository) CheckIsNotCanceled(ctx context.Context, payload financialtransactionrequest.CancelFinancialTransactionRequest) (bool, error) {
	var isNotCanceled bool

	if payload.Id != nil && payload.IdempotencyKey == nil {
		if err := f.db.QueryRow(
			ctx,
			`
				SELECT
					CASE
						WHEN canceled_at IS NULL THEN true
						ELSE false
					END AS is_not_canceled
				FROM	
					financial_transactions
				WHERE
					id = $1
			`,
			*payload.Id,
		).Scan(&isNotCanceled); err != nil {
			return false, err
		}
	}

	if payload.Id == nil && payload.IdempotencyKey != nil {
		if err := f.db.QueryRow(
			ctx,
			`
				SELECT
					CASE
						WHEN canceled_at IS NULL THEN true
						ELSE false
					END AS is_not_canceled
				FROM	
					financial_transactions ft
				WHERE
					idempotency_key = $1
			`,
			*payload.IdempotencyKey,
		).Scan(&isNotCanceled); err != nil {
			return false, err
		}
	}

	return isNotCanceled, nil
}
