package financialaccountvalidator

import (
	"context"
	"finance/internal/apperrors"
	financialaccountrequest "finance/internal/http/request/financial/financial_account"
	"finance/internal/logger"
	financialmodel "finance/models/financial"
	"fmt"
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

func (v *FinancialAccountValidator) ValidatePayload(ctx context.Context, payload financialaccountrequest.FinancialAccountRequest) error {
	errors := apperrors.ValidationErrors{}

	financialAccountByName, err := v.repo.VerifyExistsFinancialAccountName(ctx, payload.Name)

	if err != nil {
		logger.General.Error.Println("Erro ao conferir se a conta financeira já existe pelo nome: ", err)
		return err
	}

	if financialAccountByName != nil {
		errors["name"] = append(errors["name"], fmt.Sprintf("A conta financeira %s já existe, ID %d.", payload.Name, financialAccountByName.Id))
	}

	return nil
}
