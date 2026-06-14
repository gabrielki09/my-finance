package financialaccountvalidator

import (
	"context"
	"finance/internal/apperrors"
	financialrequest "finance/internal/http/request/financial"
	financialmodel "finance/models/financial"
)

type FinancialAccountRepository interface {
	VerifyExistsFinancialAccountName(ctx context.Context, financialAccountName string) (*financialmodel.FinancialAccountModel, error)
}

type FinancialAccountValidator struct {
	repo FinancialAccountRepository
}

func NewFinancialAccountValidatorValidator(financialAccountRepository FinancialAccountRepository) *FinancialAccountValidator {
	return &FinancialAccountValidator{
		repo: financialAccountRepository,
	}
}

func (v *FinancialAccountValidator) ValidatePayload(ctx context.Context, payload financialrequest.FinancialAccountRequest) error {
	errors := apperrors.ValidationErrors{}

	return nil
}
