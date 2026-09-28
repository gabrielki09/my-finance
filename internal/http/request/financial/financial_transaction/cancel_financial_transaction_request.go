package financialtransactionrequest

import (
	"finance/internal/apperrors"
	"finance/internal/logger"

	"github.com/google/uuid"
)

type CancelFinancialTransactionRequest struct {
	Id             *int    `json:"id" validated:"required"`
	IdempotencyKey *string `json:"idempotency_key" validated:"required"`
}

func (c CancelFinancialTransactionRequest) ValidatePayload() apperrors.ValidationErrors {
	logger.Info("---- Vai validar o payload da movimentação financeira via request ----")

	errors := apperrors.ValidationErrors{}

	if c.Id == nil && c.IdempotencyKey == nil {
		errors["identifier"] = append(errors["identifier"], "Ao menos um identificador da transação financeira deve ser informado.")

		return errors
	}

	if *c.Id < 1 {
		errors["id"] = append(errors["id"], "O identificador da transação financeira é inválido.")
	}

	if err := uuid.Validate(*c.IdempotencyKey); err != nil && c.Id == nil {
		errors["idempotency_key"] = append(errors["idempotency_key"], "O identificador da transação financeira é inválido.")
	}

	logger.Info("---- Terminou de validar o payload da movimentação financeira, total de erros: %d ----", len(errors))
	return errors
}
