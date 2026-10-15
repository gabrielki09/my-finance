package financialtransferservice

import (
	"context"
	financialresponse "finance/internal/http/response/financial"
	financialtransfermodel "finance/models/financial_transfer"
)

type FinancialTransferRepository interface {
	GetAll(ctx context.Context, query string, args []any) ([]financialtransfermodel.FinancialTransferModel, error)
	Create(ctx context.Context, payload any) (financialtransfermodel.FinancialTransferModel, error)
	FindByID(ctx context.Context, financialTransferID int) (financialtransfermodel.FinancialTransferModel, error)
	Delete(ctx context.Context, financialTransferID int) error
}

type FinancialTransferService struct {
	repository FinancialTransferRepository
}

func NewFinancialTransferService(repository FinancialTransferRepository) *FinancialTransferService {
	return &FinancialTransferService{
		repository: repository,
	}
}

func (c *FinancialTransferService) GetAll(ctx context.Context, query string, args []any) ([]financialresponse.FinancialTransferResponse, error) {

	return nil, nil
}

func (c *FinancialTransferService) Create(ctx context.Context, payload any) (financialresponse.FinancialTransferResponse, error) {
	return financialresponse.FinancialTransferResponse{}, nil
}

func (c *FinancialTransferService) FindByID(ctx context.Context, financialTransferID int) (financialresponse.FinancialTransferResponse, error) {
	return financialresponse.FinancialTransferResponse{}, nil
}

func (c *FinancialTransferService) Delete(ctx context.Context, financialTransferID int) error {
	return nil
}
