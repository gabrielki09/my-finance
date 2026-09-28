package financialtransfermodel

import "time"

type FinancialTransferModel struct {
	ID                   int
	TransferDate         time.Time
	SourceAccountID      int
	DestinationAccountID int
	IdempotencyKey       string
	Amount               float64
	CreatedAt            time.Time
	CanceledAt           *time.Time
}
