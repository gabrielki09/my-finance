package financialrequest

import (
	"finance/internal/apperrors"
	mxl "finance/internal/constants/max_len"
	"finance/internal/logger"
	financialmodel "finance/models/financial"
	"fmt"
	"strings"
	"time"
)

type FinancialAccountRequest struct {
	Name           string  `json:"name" validate:"required"`
	Type           string  `json:"type" validate:"required"`
	InitialBalance float64 `json:"initial_balance"`
	OpenedAt       string  `json:"opened_at" validate:"required"`
}

func validateType(t financialmodel.FinancialAccountType) bool {
	switch t {
	case financialmodel.CASH,
		financialmodel.CHECKING,
		financialmodel.SAVINGS,
		financialmodel.DIGITAL:

		return true

	default:
		return false
	}
}

func (f FinancialAccountRequest) ValidatePayload() apperrors.ValidationErrors {
	logger.General.Info.Println("---- Vai validar o payload da conta financeira via request ----")

	errors := apperrors.ValidationErrors{}

	name := strings.TrimSpace(f.Name)
	accountType := financialmodel.FinancialAccountType(f.Type)

	if name == "" {
		errors["name"] = append(errors["name"], "O nome da conta financeira é obrigatório.")
	} else if len(name) > mxl.MAX_LEN_100 {
		errors["name"] = append(errors["name"], fmt.Sprintf("O nome da conta financeira deve ter no máximo %d caracteres.", mxl.MAX_LEN_100))
	}

	if accountType == "" {
		errors["type"] = append(errors["type"], "O tipo da conta financeira é obrigatório.")
	} else if !validateType(accountType) {
		errors["type"] = append(errors["type"], "O tipo da conta financeira é inválido.")
	}

	if f.OpenedAt == "" {
		errors["opened_at"] = append(errors["opened_at"], "A data de abertura da conta financeira é obrigatória.")
	} else if _, err := time.Parse("2006-01-02", f.OpenedAt); err != nil {
		errors["opened_at"] = append(errors["opened_at"], "A data de abertura deve estar no formato YYYY-MM-DD.")
	}

	logger.General.Info.Printf("---- Terminou de validar o payload da conta financeira, total de erros: %d ----", len(errors))
	return errors
}
