package financialobligationcontroller

import (
	"context"
	"encoding/json"
	"errors"
	"finance/internal/apperrors"
	"finance/internal/helpers/response"
	"finance/internal/http/httpx"
	financialobligationrequest "finance/internal/http/request/financial/financial_obligation"
	financialresponse "finance/internal/http/response/financial"
	"finance/internal/logger"
	"net/http"
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

func (f *FinancialObligationController) GetAll(w http.ResponseWriter, r *http.Request) {
	financialObligations, err := f.service.GetAll(r.Context())

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao retornar todas as obrigações financeiras",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Todas as obrigações financeiras",
		map[string]any{"financial_obligations": financialObligations},
	))
}

func (f *FinancialObligationController) Create(w http.ResponseWriter, r *http.Request) {
	logger.General.Info.Println("FinancialObligationController - Create")

	var payload financialobligationrequest.FinancialObligationRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		logger.General.Error.Println("Erro:", err)

		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler os dados",
			httpx.DecodeErrorMessage(err),
		))
		return
	}

	financialObligation, err := f.service.Create(r.Context(), payload)

	if err != nil {
		var validationErr *apperrors.ValidationError

		if errors.As(err, &validationErr) {
			response.WriteJSON(w, http.StatusUnprocessableEntity, response.ErrorResponse(
				"Erro de validação",
				validationErr.Errors,
			))
			return
		}

		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao cadastrar a obrigação financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusCreated, response.SuccessResponse(
		"Obrigações financeira cadadastrada com sucesso",
		map[string]any{"financial_obligation": financialObligation},
	))

}
