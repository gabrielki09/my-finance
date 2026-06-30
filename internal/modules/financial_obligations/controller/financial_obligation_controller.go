package financialobligationcontroller

import (
	"context"
	"encoding/json"
	"errors"
	"finance/internal/apperrors"

	"finance/internal/helpers/getidpath"
	"finance/internal/helpers/response"
	financialobligationfiltersv1 "finance/internal/http/filter_V1/financial_obligation"
	"finance/internal/http/httpx"
	financialobligationrequest "finance/internal/http/request/financial/financial_obligation"
	financialresponse "finance/internal/http/response/financial"
	"finance/internal/logger"
	"net/http"
)

type FinancialObligationService interface {
	GetAll(context.Context, string, []any) ([]financialresponse.FinancialObligationResponse, error)
	Create(context.Context, financialobligationrequest.FinancialObligationRequest) (financialresponse.FinancialObligationResponse, error)
	Update(context.Context, financialobligationrequest.FinancialObligationRequest, int) (financialresponse.FinancialObligationResponse, error)
	Pay(context.Context, financialobligationrequest.PayFinancialObligationRequest) error
	Cancel(context.Context, int) error
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
	filters, args, err := financialobligationfiltersv1.ParseFinancialObligationFilters(r)

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
			"Erro ao conferir os filtros das obrigações financeiras",
			map[string]any{"error": err.Error()},
		))
		return
	}

	financialObligations, err := f.service.GetAll(r.Context(), filters, args)

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

func (f *FinancialObligationController) Update(w http.ResponseWriter, r *http.Request) {
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

	financialObligationId, err := getidpath.GetIdPath(r)

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler o identificador da obrigação financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	financialObligation, err := f.service.Update(r.Context(), payload, financialObligationId)

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
			"Erro ao alterar a obrigação financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Obrigações financeira alterada com sucesso",
		map[string]any{"financial_obligation": financialObligation},
	))
}

// Cancel(context.Context, int) error
func (f *FinancialObligationController) Cancel(w http.ResponseWriter, r *http.Request) {
	financialObligationId, err := getidpath.GetIdPath(r)

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler o identificador da obrigação financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	if err := f.service.Cancel(r.Context(), financialObligationId); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao cancelar a obrigação financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Obrigações financeira cancelada com sucesso",
		map[string]any{},
	))
}

func (f *FinancialObligationController) Pay(w http.ResponseWriter, r *http.Request) {
	var payload financialobligationrequest.PayFinancialObligationRequest

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

	if err := f.service.Pay(r.Context(), payload); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao pagar a obrigação financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Obrigações financeira paga com sucesso",
		map[string]any{},
	))
}
