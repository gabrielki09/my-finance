package financialmodel

import "time"

type FinancialTransactionsMovementType string

const (
	ENTRY FinancialTransactionsMovementType = "entry"
	EXIT  FinancialTransactionsMovementType = "exit"
)

type FinancialTransactionsOperationType string

const (
	ORIGINAL   FinancialTransactionsOperationType = "original"
	ADJUSTMENT FinancialTransactionsOperationType = "adjustment"
	REVERSAL   FinancialTransactionsOperationType = "reversal"
)

type FinancialTransactionsModel struct {
	Id                    int
	FinancialAccountId    int
	CategoryId            int
	ReversedTransactionId int
	Description           string
	MovementType          FinancialTransactionsMovementType
	OperationType         FinancialTransactionsOperationType
	Amount                float64
	MovementDate          time.Time
	ReferenceDate         time.Time
	OriginType            string
	OriginId              int
	Idempotency_key       string
	CreatedAt             time.Time
	CanceledAt            *time.Time
}
