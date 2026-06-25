package financialobligationservice

import (
	"context"
	"finance/internal/apperrors"
	financialobligationrequest "finance/internal/http/request/financial/financial_obligation"
	financialresponse "finance/internal/http/response/financial"
	"finance/internal/logger"
	financialmapper "finance/internal/mapper/financial"
	financialmodel "finance/models/financial"
)

type FinancialObligationRepository interface {
	GetAll(ctx context.Context, query string, args []any) ([]financialmodel.FinancialObligationModel, error)
	Create(ctx context.Context, payload financialobligationrequest.FinancialObligationRequest) (financialmodel.FinancialObligationModel, error)
	Update(ctx context.Context, payload financialobligationrequest.FinancialObligationRequest, financialObligationId int) (financialmodel.FinancialObligationModel, error)
	Cancel(ctx context.Context, financialObligationId int) error
	FindById(ctx context.Context, financialObligationId int) (financialmodel.FinancialObligationModel, error)
}

type FinancialObligationService struct {
	repository FinancialObligationRepository
	request    financialobligationrequest.FinancialObligationRequest
}

func NewFinancialObligationService(repository FinancialObligationRepository, financialobligationrequest *financialobligationrequest.FinancialObligationRequest) *FinancialObligationService {
	return &FinancialObligationService{
		repository: repository,
		request:    *financialobligationrequest,
	}
}

func (f *FinancialObligationService) GetAll(ctx context.Context, query string, args []any) ([]financialresponse.FinancialObligationResponse, error) {

	financialObligations, err := f.repository.GetAll(ctx, query, args)

	if err != nil {
		logger.General.Error.Println("Erro ao retornar todas as obrigações financeiras:", err)
		return []financialresponse.FinancialObligationResponse{}, err
	}

	return financialmapper.ToFinancialObligationResponseList(financialObligations), nil

}

func (f *FinancialObligationService) Create(ctx context.Context, payload financialobligationrequest.FinancialObligationRequest) (financialresponse.FinancialObligationResponse, error) {

	validation := payload.ValidatePayload()

	if len(validation) > 0 {
		return financialresponse.FinancialObligationResponse{}, apperrors.NewValidationError(validation)
	}

	financialObligation, err := f.repository.Create(ctx, payload)

	if err != nil {
		logger.General.Error.Println("Erro ao criar a conta financeira:", err)
		return financialresponse.FinancialObligationResponse{}, err
	}

	return financialmapper.ToFinancialObligationResponse(financialObligation), nil
}

func (f *FinancialObligationService) Update(ctx context.Context, payload financialobligationrequest.FinancialObligationRequest, financialObligationId int) (financialresponse.FinancialObligationResponse, error) {

	logger.General.Info.Println("FinancialObligationService - Update")

	validation := payload.ValidatePayload()

	if len(validation) > 0 {
		return financialresponse.FinancialObligationResponse{}, apperrors.NewValidationError(validation)
	}

	financialObligation, err := f.repository.Update(ctx, payload, financialObligationId)

	if err != nil {
		logger.General.Error.Println("Erro ao criar a conta financeira:", err)
		return financialresponse.FinancialObligationResponse{}, err
	}

	return financialmapper.ToFinancialObligationResponse(financialObligation), nil
}

func (f *FinancialObligationService) Cancel(ctx context.Context, financialObligationId int) error {

	if err := f.request.ValidateCancel(ctx, financialObligationId); err != nil {
		logger.General.Error.Println("Erro ao validar a obrigação financeira para o cancelamento:", err)
		return err
	}

	if err := f.repository.Cancel(ctx, financialObligationId); err != nil {
		logger.General.Error.Println("Erro ao cancelar a obrigação financeira:", err)
		return err
	}

	return nil
}
