package financialresponse

import "time"

type FinancialTransferResponse struct {
	ID                   int        `json:"id"`
	TransferDate         time.Time  `json:"transfer_date"`
	SourceAccountID      int        `json:"source_account_id"`
	DestinationAccountID int        `json:"destination_account_id"`
	IdempotencyKey       string     `json:"idempotency_key"`
	Amount               float64    `json:"amount"`
	CreatedAt            time.Time  `json:"created_at"`
	CanceledAt           *time.Time `json:"canceled_at,omitempty"`
}
