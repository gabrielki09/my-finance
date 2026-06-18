package financialtransactioncontroller

import (
	"context"
	"encoding/json"
	"errors"
	"finance/internal/apperrors"
	"finance/internal/helpers/getidpath"
	"finance/internal/helpers/response"
	"finance/internal/http/httpx"
	financialtransactionrequest "finance/internal/http/request/financial/financial_transaction"
	financialresponse "finance/internal/http/response/financial"
	"finance/internal/logger"
	"net/http"

	"github.com/google/uuid"
)

type FinancialTransactionService interface {
	GetAll(ctx context.Context) ([]financialresponse.FinancialTransactionsResponse, error)
	FindById(ctx context.Context, financialTransactionId int) (financialresponse.FinancialTransactionsResponse, error)
	FindByKey(ctx context.Context, key string) (financialresponse.FinancialTransactionsResponse, error)
	CreateMovement(ctx context.Context, payload financialtransactionrequest.FinancialTransactionRequest) (financialresponse.FinancialTransactionsResponse, error)
	CancelMovement(ctx context.Context, payload financialtransactionrequest.CancelFinancialTransactionRequest) error
}

type FinancialTransactionController struct {
	service FinancialTransactionService
}

func NewFinancialTransactionController(service FinancialTransactionService) *FinancialTransactionController {
	return &FinancialTransactionController{
		service: service,
	}
}

func (c *FinancialTransactionController) GetAll(w http.ResponseWriter, r *http.Request) {
	logger.General.Info.Println("FinancialTransactionController - GetAll")

	financialTransactions, err := c.service.GetAll(r.Context())

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao retornar todas as transações financeiras",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Todas as transações financeiras",
		map[string]any{"financial_transactions": financialTransactions},
	))
}

func (c *FinancialTransactionController) FindById(w http.ResponseWriter, r *http.Request) {
	logger.General.Info.Println("FinancialTransactionController - FindById")

	financialTransactionId, err := getidpath.GetIdPath(r)

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler o identificador da transação financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	financialTransaction, err := c.service.FindById(r.Context(), financialTransactionId)

	if err != nil {
		if err == apperrors.ErrNotFound {
			response.WriteJSON(w, http.StatusNotFound, response.ErrorResponse(
				"Transação financeira não localizada",
				map[string]any{},
			))
			return
		}

		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao localizar a transação financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Transação financeira localizada",
		map[string]any{"financial_transaction": financialTransaction},
	))
}

func (c *FinancialTransactionController) FindByKey(w http.ResponseWriter, r *http.Request) {
	logger.General.Info.Println("FinancialTransactionController - FindByKey")

	financialTransactionKey := r.PathValue("key")

	if err := uuid.Validate(financialTransactionKey); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Identificador da transação financeira fora do formato esperado.",
			map[string]any{"error": err.Error()},
		))
		return
	}

	financialTransaction, err := c.service.FindByKey(r.Context(), financialTransactionKey)

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler o identificador da transação financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Todas as transação financeira",
		map[string]any{"financial_transaction": financialTransaction},
	))
}

func (c *FinancialTransactionController) CreateMovement(w http.ResponseWriter, r *http.Request) {
	logger.General.Info.Println("FinancialTransactionController - CreateMovement")

	var payload financialtransactionrequest.FinancialTransactionRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler os dados",
			httpx.DecodeErrorMessage(err),
		))
		return
	}

	financialTransaction, err := c.service.CreateMovement(r.Context(), payload)

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
			"Erro ao cadastrar a transação financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusCreated, response.SuccessResponse(
		"Transação financeira cadastrada com sucesso!",
		map[string]any{"financial_transaction": financialTransaction},
	))
}

func (f *FinancialTransactionController) CancelMovement(w http.ResponseWriter, r *http.Request) {
	var payload financialtransactionrequest.CancelFinancialTransactionRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler os dados",
			httpx.DecodeErrorMessage(err),
		))
		return
	}

	if err := f.service.CancelMovement(r.Context(), payload); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao cancelar a transação financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Transação financeira cancelada com sucesso!",
		map[string]any{},
	))
}
