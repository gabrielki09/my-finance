package financialobligationservice

import (
	"context"
	"finance/internal/apperrors"
	financialobligationrequest "finance/internal/http/request/financial/financial_obligation"
	financialresponse "finance/internal/http/response/financial"
	"finance/internal/logger"
	financialmapper "finance/internal/mapper/financial"
	financialobligationvalidator "finance/internal/modules/financial_obligations/validator"
	financialmodel "finance/models/financial"
)

type FinancialObligationRepository interface {
	GetAll(ctx context.Context) ([]financialmodel.FinancialObligationModel, error)
	Create(ctx context.Context, payload financialobligationrequest.FinancialObligationRequest) (financialmodel.FinancialObligationModel, error)
}

type FinancialObligationService struct {
	repository FinancialObligationRepository
	validator  *financialobligationvalidator.FinancialObligationValidator
}

func NewFinancialObligationService(repository FinancialObligationRepository, validator *financialobligationvalidator.FinancialObligationValidator) *FinancialObligationService {
	return &FinancialObligationService{
		repository: repository,
		validator:  validator,
	}
}

func (f *FinancialObligationService) GetAll(ctx context.Context) ([]financialresponse.FinancialObligationResponse, error) {

	financialObligations, err := f.repository.GetAll(ctx)

	if err != nil {
		logger.General.Error.Println("Erro ao criar a conta financeira:", err)
		return []financialresponse.FinancialObligationResponse{}, err
	}

	return financialmapper.ToFinancialObligationResponseList(financialObligations), nil

}
func (f *FinancialObligationService) Create(ctx context.Context, payload financialobligationrequest.FinancialObligationRequest) (financialresponse.FinancialObligationResponse, error) {

	validation := payload.ValidatePayload()

	if len(validation) > 0 {
		return financialresponse.FinancialObligationResponse{}, apperrors.NewValidationError(validation)
	}

	if err := f.validator.ValidatePayload(ctx, payload); err != nil {
		logger.General.Error.Println("Erro na validação de dados:", err)
		return financialresponse.FinancialObligationResponse{}, err
	}

	financialObligation, err := f.repository.Create(ctx, payload)

	if err != nil {
		logger.General.Error.Println("Erro ao criar a conta financeira:", err)
		return financialresponse.FinancialObligationResponse{}, err
	}

	return financialmapper.ToFinancialObligationResponse(financialObligation), nil
}
