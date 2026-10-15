package financialresponse

import (
	financialmodel "finance/models/financial"
	"time"
)

type FinancialTransactionsResponse struct {
	Id                    int                                               `json:"id"`
	FinancialAccountId    int                                               `json:"financial_account_id"`
	CategoryId            int                                               `json:"category_id"`
	ReversedTransactionId *int                                              `json:"reversed_transaction_id"`
	Description           string                                            `json:"description"`
	MovementType          financialmodel.FinancialTransactionsMovementType  `json:"movement_type"`
	OperationType         financialmodel.FinancialTransactionsOperationType `json:"operation_type"`
	Amount                float64                                           `json:"amount"`
	MovementDate          time.Time                                         `json:"movement_date"`
	ReferenceDate         time.Time                                         `json:"reference_date"`
	OriginType            *string                                           `json:"origin_type"`
	OriginId              *int                                              `json:"origin_id"`
	IdempotencyKey        string                                            `json:"idempotency_key"`
	CreatedAt             time.Time                                         `json:"created_at"`
	CanceledAt            *time.Time                                        `json:"canceled_at"`
}
