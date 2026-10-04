package financialtransfercontroller

import (
	"context"
	"encoding/json"
	"errors"
	"finance/internal/apperrors"
	"finance/internal/helpers/response"
	financialtransferfiltersv1 "finance/internal/http/filter_v1/financial_transfer"
	"finance/internal/http/httpx"
	financialresponse "finance/internal/http/response/financial"
	"finance/internal/logger"

	"net/http"
)

type FinancialTransferService interface {
	GetAll(context.Context, string, []any) ([]financialresponse.FinancialTransferResponse, error)
	Create(context.Context, any) (financialresponse.FinancialTransferResponse, error)
	FindByID(context.Context, int) (financialresponse.FinancialTransferResponse, error)
	Delete(context.Context, int) error
}

type FinancialTransferController struct {
	service FinancialTransferService
}

func NewFinancialTransferController(service FinancialTransferService) *FinancialTransferController {
	return &FinancialTransferController{
		service: service,
	}
}

func (f *FinancialTransferController) GetAll(w http.ResponseWriter, r *http.Request) {
	filters, args, err := financialtransferfiltersv1.ParseFinancialTransferFilters(r)
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
			"Erro ao conferir os filtras das transferências financeiras",
			map[string]any{"error": err.Error()},
		))
		return
	}

	financialTransfers, err := f.service.GetAll(r.Context(), filters, args)

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao retornar todas as transferências financeiras",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Todas as transferênicas financeiras",
		map[string]any{"financial_transfers": financialTransfers},
	))
}

func (f *FinancialTransferController) Create(w http.ResponseWriter, r *http.Request) {
	logger.Info("FinancialTransferController - Create")

	var payload any

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		logger.Error("Erro ao conferir os dados do payload:", err)

		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler os dados",
			httpx.DecodeErrorMessage(err),
		))
		return
	}

	financialTransfer, err := f.service.Create(r.Context(), payload)
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
			"Erro ao cadastrar a transferência financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusCreated, response.SuccessResponse(
		"Transferênica cadadastrada com sucesso",
		map[string]any{"financial_transfer": financialTransfer},
	))
}

func (f *FinancialTransferController) FindByID(w http.ResponseWriter, r *http.Request) {
}

func (f *FinancialTransferController) Delete(w http.ResponseWriter, r *http.Request) {
}
