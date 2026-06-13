package financialaccountcontroller

import (
	"context"
	"encoding/json"
	"errors"
	"finance/internal/apperrors"
	"finance/internal/helpers/getidpath"
	"finance/internal/helpers/response"
	"finance/internal/http/httpx"
	financialrequest "finance/internal/http/request/financial"
	financialresponse "finance/internal/http/response/financial"
	"finance/internal/logger"
	"net/http"
)

type FinancialAccountService interface {
	GetAll(context.Context) ([]financialresponse.FinancialAccountResponse, error)
	Create(context.Context, financialrequest.FinancialAccountRequest) (financialresponse.FinancialAccountResponse, error)
	Update(context.Context, financialrequest.FinancialAccountRequest, int) (financialresponse.FinancialAccountResponse, error)
	FindById(context.Context, int) (financialresponse.FinancialAccountResponse, error)
	Delete(context.Context, int) error
	Active(context.Context, int) error
}

type FinancialAccountController struct {
	service FinancialAccountService
}

func NewFinancialAccountController(service FinancialAccountService) *FinancialAccountController {
	return &FinancialAccountController{
		service: service,
	}
}

func (f *FinancialAccountController) GetAll(w http.ResponseWriter, r *http.Request) {
	financialAccounts, err := f.service.GetAll(r.Context())

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao retornar todas as contas financeiras",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Todas as contas financeiras",
		map[string]any{"financial_accounts": financialAccounts},
	))
}

func (f *FinancialAccountController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Body == http.NoBody {
		logger.General.Error.Println("Payload vazio")

		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Dados ausentes",
			map[string]any{},
		))
		return
	}

	var payload financialrequest.FinancialAccountRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler os dados",
			httpx.DecodeErrorMessage(err),
		))
		return
	}

	financialAccount, err := f.service.Create(r.Context(), payload)

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
			"Erro ao cadastrar a conta financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusCreated, response.SuccessResponse(
		"Conta financeira cadastrada com sucesso!",
		map[string]any{"financial_account": financialAccount},
	))
}

func (f *FinancialAccountController) Update(w http.ResponseWriter, r *http.Request) {
	if r.Body == http.NoBody {
		logger.General.Error.Println("Payload vazio")

		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Dados ausentes",
			map[string]any{},
		))
		return
	}

	var payLoad financialrequest.FinancialAccountRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payLoad); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler os dados",
			map[string]any{"error": err.Error()},
		))
		return
	}

	id, err := getidpath.GetIdPath(r)

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler o identificador da conta financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	financialAccount, err := f.service.Update(r.Context(), payLoad, id)

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
			"Erro ao alterar a conta financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusCreated, response.SuccessResponse(
		"Conta financeira alterada com sucesso!",
		map[string]any{"financial_account": financialAccount},
	))
}

func (f *FinancialAccountController) FindById(w http.ResponseWriter, r *http.Request) {
	id, err := getidpath.GetIdPath(r)

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler o identificador da conta financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	financialAccount, err := f.service.FindById(r.Context(), id)

	if err != nil {
		if err == apperrors.ErrNotFound {
			response.WriteJSON(w, http.StatusNotFound, response.ErrorResponse(
				"Conta financeira não localizada",
				map[string]any{},
			))
			return
		}

		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao localizar a conta financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Conta financeira localizada com sucesso!",
		map[string]any{"financial_account": financialAccount},
	))
}

func (f *FinancialAccountController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getidpath.GetIdPath(r)

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler o identificador da conta financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	if err := f.service.Delete(r.Context(), id); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao desativar a conta financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Conta financeira desativada com sucesso!",
		map[string]any{},
	))
}

func (f *FinancialAccountController) Active(w http.ResponseWriter, r *http.Request) {
	id, err := getidpath.GetIdPath(r)

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler o identificador da conta financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	if err := f.service.Active(r.Context(), id); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ativar a conta financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Conta financeira ativada com sucesso!",
		map[string]any{},
	))
}
