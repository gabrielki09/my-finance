package financialobligationcontroller

import (
	"context"
	financialobligationrequest "finance/internal/http/request/financial/financial_obligation"
	financialresponse "finance/internal/http/response/financial"
)

type FinancialObligationService interface {
	GetAll(context.Context) ([]financialresponse.FinancialObligationResponse, error)
	Create(context.Context, financialobligationrequest.FinancialObligationRequest) (financialresponse.FinancialObligationResponse, error)
}

type FinancialObligationController struct {
	service FinancialObligationService
}

func NewFinancialObligationController(service FinancialObligationService) *FinancialObligationController {
	return &FinancialObligationController{
		service: service,
	}
}
