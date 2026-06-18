package financialtransactionvalidator

import (
	"context"
	"finance/internal/apperrors"
	financialtransactionrequest "finance/internal/http/request/financial/financial_transaction"
	"finance/internal/logger"
	financialmodel "finance/models/financial"
)

type FinancialTransactionRepository interface {
	ValidCategoryType(ctx context.Context, categoryId int, movementType financialmodel.FinancialTransactionsMovementType) (bool, error)
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

	errors := apperrors.ValidationErrors{}

	isInvalidType, err := f.financialTransactionRepository.ValidCategoryType(ctx, payload.CategoryId, payload.MovementType)

	if err != nil {
		logger.General.Error.Println("Erro ao válidar se o tipo da categoria é coerente com o tipo de movimento financeiro:", err)
		return err
	}

	if isInvalidType {
		errors["category_id"] = append(errors["category_id"], "Tipo da categoria incoerente com o tipo da movimentação financeira")
	}

	logger.General.Info.Printf("---- Terminou de validar o payload da categoria, total de erros: %d ----", len(errors))

	if len(errors) > 0 {
		return apperrors.NewValidationError(errors)
	}

	return nil
}
