package financialtransactionvalidator

import (
	"context"
	"errors"
	"finance/internal/apperrors"
	financialtransactionrequest "finance/internal/http/request/financial/financial_transaction"
	"finance/internal/logger"
	financialmodel "finance/models/financial"

	"github.com/jackc/pgx/v5"
)

type FinancialTransactionRepository interface {
	ValidCategoryType(ctx context.Context, categoryId int, movementType financialmodel.FinancialTransactionsMovementType) (bool, error)
	ValidateIsSameIdAndIdempotencyKey(ctx context.Context, financialTransactionId int, idempotencyKey string) (bool, error)
}

type FinancialTransactionValidator struct {
	financialTransactionRepository FinancialTransactionRepository
}

func NewFinancialTransactionValidator(financialTransactionRepository FinancialTransactionRepository) *FinancialTransactionValidator {
	return &FinancialTransactionValidator{
		financialTransactionRepository: financialTransactionRepository,
	}
}

func (f *FinancialTransactionValidator) ValidatePayload(ctx context.Context, payload financialtransactionrequest.FinancialTransactionRequest) error {
	logger.General.Info.Println("---- Vai validar o payload do movimento financeiro via db ----")

	appErrors := apperrors.ValidationErrors{}

	isInvalidType, err := f.financialTransactionRepository.ValidCategoryType(ctx, payload.CategoryId, payload.MovementType)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			appErrors["category_id"] = append(appErrors["category_id"], "Categoria não localizada.")
			return apperrors.NewValidationError(appErrors)
		}

		logger.General.Error.Println("Erro ao válidar se o tipo da transação financeira é coerente com o tipo de movimento financeiro:", err)
		return err
	}

	if isInvalidType {
		appErrors["category_id"] = append(appErrors["category_id"], "Tipo da categoria incoerente com o tipo da movimentação financeira.")
	}

	logger.General.Info.Printf("---- Terminou de validar o payload da transação financeira, total de erros: %d ----", len(appErrors))

	if len(appErrors) > 0 {
		return apperrors.NewValidationError(appErrors)
	}

	return nil
}
