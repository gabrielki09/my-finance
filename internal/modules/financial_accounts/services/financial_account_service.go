package financialaccountservice

import (
	"context"
	"finance/internal/apperrors"
	financialrequest "finance/internal/http/request/financial"
	financialresponse "finance/internal/http/response/financial"
	"finance/internal/logger"
	financialmapper "finance/internal/mapper/financial"
	financialaccountvalidator "finance/internal/modules/financial_accounts/validator"
	financialmodel "finance/models/financial"
)

type FinancialAccountRepository interface {
	GetAll(r context.Context) ([]financialmodel.FinancialAccountModel, error)
	Create(r context.Context, payload financialrequest.FinancialAccountRequest) (financialmodel.FinancialAccountModel, error)
	Update(r context.Context, payload financialrequest.FinancialAccountRequest, financialAccountId int) (financialmodel.FinancialAccountModel, error)
	FindById(r context.Context, financialAccountId int) (financialmodel.FinancialAccountModel, error)
	Delete(r context.Context, financialAccountId int) error
	Active(r context.Context, financialAccountId int) error
}

type FinancialAccountService struct {
	repository FinancialAccountRepository
	validator  *financialaccountvalidator.FinancialAccountValidator
}

func NewFinancialAccountService(repository FinancialAccountRepository, validator *financialaccountvalidator.FinancialAccountValidator) *FinancialAccountService {
	return &FinancialAccountService{
		repository: repository,
		validator:  validator,
	}
}

func (f *FinancialAccountService) GetAll(ctx context.Context) ([]financialresponse.FinancialAccountResponse, error) {
	financialAccounts, err := f.repository.GetAll(ctx)

	if err != nil {
		logger.General.Error.Println("Erro ao consultar todas as conta financeiras:", err)
		return nil, err
	}

	return financialmapper.ToFinancialAccountResponseList(financialAccounts), nil
}

func (f *FinancialAccountService) Create(ctx context.Context, payload financialrequest.FinancialAccountRequest) (financialresponse.FinancialAccountResponse, error) {

	validation := payload.ValidatePayload()

	if len(validation) > 0 {
		return financialresponse.FinancialAccountResponse{}, apperrors.NewValidationError(validation)
	}

	if err := f.validator.ValidatePayload(ctx, payload); err != nil {
		return financialresponse.FinancialAccountResponse{}, err
	}

	financialAccount, err := f.repository.Create(ctx, payload)

	if err != nil {
		logger.General.Error.Println("Erro ao criar a conta financeira:", err)
		return financialresponse.FinancialAccountResponse{}, err
	}

	return financialmapper.ToFinancialAccountResponse(financialAccount), nil
}

func (f *FinancialAccountService) Update(ctx context.Context, payload financialrequest.FinancialAccountRequest, financialAccountId int) (financialresponse.FinancialAccountResponse, error) {

	validation := payload.ValidatePayload()

	if len(validation) > 0 {
		return financialresponse.FinancialAccountResponse{}, apperrors.NewValidationError(validation)
	}

	if err := f.validator.ValidatePayload(ctx, payload); err != nil {
		return financialresponse.FinancialAccountResponse{}, err
	}

	financialAccount, err := f.repository.Update(ctx, payload, financialAccountId)

	if err != nil {
		logger.General.Error.Println("Erro ao alterar a conta financeira:", err)
		return financialresponse.FinancialAccountResponse{}, err
	}

	return financialmapper.ToFinancialAccountResponse(financialAccount), nil
}

func (f *FinancialAccountService) FindById(ctx context.Context, financialAccountId int) (financialresponse.FinancialAccountResponse, error) {
	financialAccount, err := f.repository.FindById(ctx, financialAccountId)

	if err != nil {
		logger.General.Error.Println("Erro ao localizar a conta financeira:", err)
		return financialresponse.FinancialAccountResponse{}, err
	}

	return financialmapper.ToFinancialAccountResponse(financialAccount), nil
}

func (f *FinancialAccountService) Delete(ctx context.Context, financialAccountId int) error {
	if err := f.repository.Delete(ctx, financialAccountId); err != nil {
		logger.General.Error.Println("Erro ao deletar a conta financeira:", err)
		return err
	}

	return nil
}

func (f *FinancialAccountService) Active(ctx context.Context, financialAccountId int) error {
	if err := f.repository.Active(ctx, financialAccountId); err != nil {
		logger.General.Error.Println("Erro ao ativar a conta financeira:", err)
		return err
	}

	return nil
}
