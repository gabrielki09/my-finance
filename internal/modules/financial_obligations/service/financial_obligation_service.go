package financialobligationservice

import (
	"context"
	financialobligationrequest "finance/internal/http/request/financial/financial_obligation"
	financialresponse "finance/internal/http/response/financial"
	financialmodel "finance/models/financial"
)

type FinancialObligationRepository interface {
	GetAll(ctx context.Context) ([]financialresponse.FinancialObligationResponse, error)
	Create(ctx context.Context, payload financialobligationrequest.FinancialObligationRequest) (financialmodel.FinancialObligationModel, error)
}

type FinancialObligationService struct {
	repository FinancialObligationRepository
}

func NewFinancialObligationService(repository FinancialObligationRepository) *FinancialObligationService {
	return &FinancialObligationService{
		repository: repository,
	}
}
