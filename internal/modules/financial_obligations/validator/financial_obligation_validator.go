package financialobligationvalidator

import (
	"context"
	"finance/internal/apperrors"
	financialobligationrequest "finance/internal/http/request/financial/financial_obligation"
	"finance/internal/logger"
	financialmodel "finance/models/financial"
)

type FinancialObligationRepository interface {
	ValidCategoryType(ctx context.Context, categoryId int, movementType financialmodel.FinancialObligationsTypes) (bool, error)
}

type FinancialObligationValidator struct {
	repo FinancialObligationRepository
}

func NewFinancialObligationValidatorValidator(repo FinancialObligationRepository) *FinancialObligationValidator {
	return &FinancialObligationValidator{
		repo: repo,
	}
}

func (f *FinancialObligationValidator) ValidatePayload(ctx context.Context, payload financialobligationrequest.FinancialObligationRequest) error {

	logger.General.Info.Println("---- Vai validar o payload do movimento financeiro via db ----")

	errors := apperrors.ValidationErrors{}

	isInvalidType, err := f.repo.ValidCategoryType(ctx, payload.CategoryId, payload.Type)

	if err != nil {
		logger.General.Error.Println("Erro ao válidar se o tipo da categoria é coerente com o tipo de movimento financeiro:", err)
		return err
	}

	if isInvalidType {
		errors["category_id"] = append(errors["category_id"], "Tipo da categoria incoerente com o tipo da obrigação financeira")
	}

	logger.General.Info.Printf("---- Terminou de validar o payload da transação financeira, total de erros: %d ----", len(errors))

	if len(errors) > 0 {
		return apperrors.NewValidationError(errors)
	}

	return nil
}
