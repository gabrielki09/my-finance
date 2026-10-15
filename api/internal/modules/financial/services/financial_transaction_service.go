package financialtransactionservice

import (
	"context"
	"finance/internal/apperrors"
	financialtransactionrequest "finance/internal/http/request/financial/financial_transaction"
	financialresponse "finance/internal/http/response/financial"
	"finance/internal/logger"
	financialmapper "finance/internal/mapper/financial"
	financialtransactionvalidator "finance/internal/modules/financial/validator"
	financialmodel "finance/models/financial"
	"fmt"
)

type FinancialTransactionRepository interface {
	GetAll(ctx context.Context) ([]financialmodel.FinancialTransactionsModel, error)
	FindById(ctx context.Context, financialTransactionId int) (financialmodel.FinancialTransactionsModel, error)
	FindByKey(ctx context.Context, key string) (financialmodel.FinancialTransactionsModel, error)
	CreateMovement(ctx context.Context, payload financialtransactionrequest.FinancialTransactionRequest) (financialmodel.FinancialTransactionsModel, error)
	CancelMovement(ctx context.Context, payload financialtransactionrequest.CancelFinancialTransactionRequest) error

	CheckIsNotCanceled(ctx context.Context, payload financialtransactionrequest.CancelFinancialTransactionRequest) (bool, error)
}

type FinancialTransactionService struct {
	repository FinancialTransactionRepository
	validator  *financialtransactionvalidator.FinancialTransactionValidator
}

func NewFinancialTransactionService(repository FinancialTransactionRepository, validator *financialtransactionvalidator.FinancialTransactionValidator) *FinancialTransactionService {
	return &FinancialTransactionService{
		repository: repository,
		validator:  validator,
	}
}

func (f *FinancialTransactionService) GetAll(ctx context.Context) ([]financialresponse.FinancialTransactionsResponse, error) {
	financialTransactions, err := f.repository.GetAll(ctx)

	if err != nil {
		logger.Error("Erro ao retornar todas as transações financeiras:", err)
		return []financialresponse.FinancialTransactionsResponse{}, err
	}

	return financialmapper.ToFinancialTransactionResponseList(financialTransactions), nil
}

func (f *FinancialTransactionService) FindById(ctx context.Context, financialTransactionId int) (financialresponse.FinancialTransactionsResponse, error) {
	financialTransaction, err := f.repository.FindById(ctx, financialTransactionId)

	if err != nil {
		logger.Error("Erro ao retornar todas as transações financeiras:", err)
		return financialresponse.FinancialTransactionsResponse{}, err

	}

	return financialmapper.ToFinancialTransactionResponse(financialTransaction), nil
}

func (f *FinancialTransactionService) FindByKey(ctx context.Context, key string) (financialresponse.FinancialTransactionsResponse, error) {

	financialTransaction, err := f.repository.FindByKey(ctx, key)

	if err != nil {
		logger.Error("Erro ao retornar todas as transações financeiras:", err)

		return financialresponse.FinancialTransactionsResponse{}, err
	}

	return financialmapper.ToFinancialTransactionResponse(financialTransaction), nil
}

func (f *FinancialTransactionService) CreateMovement(ctx context.Context, payload financialtransactionrequest.FinancialTransactionRequest) (financialresponse.FinancialTransactionsResponse, error) {
	validation := payload.ValidatePayload()

	if len(validation) > 0 {
		return financialresponse.FinancialTransactionsResponse{}, apperrors.NewValidationError(validation)
	}

	if err := f.validator.ValidatePayload(ctx, payload); err != nil {
		return financialresponse.FinancialTransactionsResponse{}, err
	}

	financialTransaction, err := f.repository.CreateMovement(ctx, payload)

	if err != nil {
		return financialresponse.FinancialTransactionsResponse{}, err
	}

	return financialmapper.ToFinancialTransactionResponse(financialTransaction), nil
}

func (f *FinancialTransactionService) CancelMovement(ctx context.Context, payload financialtransactionrequest.CancelFinancialTransactionRequest) error {

	isNotCanceled, err := f.repository.CheckIsNotCanceled(ctx, payload)
	if err != nil {
		return err
	}

	if !isNotCanceled {
		return fmt.Errorf("Transação já cancelada.")
	}

	return nil
}
