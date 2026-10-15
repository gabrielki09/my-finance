package financialtransferrequest

import (
	"finance/internal/apperrors"
	"time"
)

type FinancialTransferRequest struct {
	TransferDate         time.Time `json:"transfer_date" validate:"required"`
	SourceAccountID      int       `json:"source_account_id" validate:"required"`
	DestinationAccountID int       `json:"destination_account_id" validate:"required"`
	IdempotencyKey       string    `json:"idempotency_key" validate:"required"`
	Amount               float64   `json:"amount" validate:"required"`
}

func (f FinancialTransferRequest) ValidatePayload() apperrors.ValidationErrors {
	errors := apperrors.ValidationErrors{}

	return errors
}
