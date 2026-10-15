package financialtransferrepository

import (
	"context"
	transactionhelper "finance/internal/helpers/transaction"
	financialtransferrequest "finance/internal/http/request/financial/financial_transfer"
	"finance/internal/logger"
	financialtransfermodel "finance/models/financial_transfer"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FinancialTransferRepository struct {
	db *pgxpool.Pool
}

func NewFinancialTransferRepository(db *pgxpool.Pool) *FinancialTransferRepository {
	return &FinancialTransferRepository{
		db: db,
	}
}

func (f *FinancialTransferRepository) GetAll(ctx context.Context, query string, args []any) ([]financialtransfermodel.FinancialTransferModel, error) {
	var financialTransfers []financialtransfermodel.FinancialTransferModel

	financialTransferRows, err := f.db.Query(
		ctx,
		query,
		args...,
	)
	if err != nil {
		logger.Error("Erro ao executar o select:", err)
		return []financialtransfermodel.FinancialTransferModel{}, err
	}

	defer financialTransferRows.Close()

	for financialTransferRows.Next() {
		var financialTransfer financialtransfermodel.FinancialTransferModel

		if err := financialTransferRows.Scan(
			&financialTransfer.ID,
			&financialTransfer.TransferDate,
			&financialTransfer.SourceAccountID,
			&financialTransfer.DestinationAccountID,
			&financialTransfer.IdempotencyKey,
			&financialTransfer.CreatedAt,
			&financialTransfer.CanceledAt,
		); err != nil {
			logger.Error("Erro ao ler os dados da consulta:", err)
			return []financialtransfermodel.FinancialTransferModel{}, err
		}

		financialTransfers = append(financialTransfers, financialTransfer)
	}

	if err := financialTransferRows.Err(); err != nil {
		logger.Error("Erro ao ler os dados da consulta:", err)
		return []financialtransfermodel.FinancialTransferModel{}, err
	}

	return financialTransfers, nil
}

func (f *FinancialTransferRepository) Create(ctx context.Context, payload financialtransferrequest.FinancialTransferRequest) (financialtransfermodel.FinancialTransferModel, error) {
	financialTransfer, err := transactionhelper.WithTransactionResult(
		ctx,
		f.db,
		func(tx pgx.Tx) (financialtransfermodel.FinancialTransferModel, error) {
			var financialTransfer financialtransfermodel.FinancialTransferModel

			if err := tx.QueryRow(
				ctx,
				`	
					INSERT INTO financial_transfer (
						transfer_date,
						source_account_id,
						destination_account_id,
						amount
					) VALUES (
						$1,
						$2,
						$3,
						$4
					) RETURNING
					 	id,
						transfer_date,
						source_account_id,
						destination_account_id,
						idempotency_key,
						amount,
						created_at					
				`,
				payload.TransferDate,
				payload.SourceAccountID,
				payload.DestinationAccountID,
				payload.Amount,
			).Scan(
				&financialTransfer.ID,
				&financialTransfer.TransferDate,
				&financialTransfer.SourceAccountID,
				&financialTransfer.DestinationAccountID,
				&financialTransfer.IdempotencyKey,
				&financialTransfer.Amount,
				&financialTransfer.CreatedAt,
			); err != nil {
				logger.Error("Erro ao criar a transferência financeira:", err)
				return financialtransfermodel.FinancialTransferModel{}, err
			}

			sourceFinancialAccountCmd, err := tx.Exec(
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
			)
			if err != nil {
				logger.Error("Erro ao criar registro de trasnferência: ", err)
				return financialtransfermodel.FinancialTransferModel{}, err
			}

			if sourceFinancialAccountCmd.RowsAffected() == 0 {
				return financialtransfermodel.FinancialTransferModel{}, fmt.Errorf("nem um registro afetado")
			}

			return financialTransfer, nil
		},
	)
	if err != nil {
		return financialtransfermodel.FinancialTransferModel{}, err
	}

	return financialTransfer, nil
}

func (f *FinancialTransferRepository) FindByID(ctx context.Context, financialTransferID int) (financialtransfermodel.FinancialTransferModel, error) {
	return financialtransfermodel.FinancialTransferModel{}, nil
}

func (f *FinancialTransferRepository) Delete(ctx context.Context, financialTransferID int) error {
	return nil
}
