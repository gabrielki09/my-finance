package financialmodel

import "time"

type FinancialObligationsStatus string

const (
	PENDING           FinancialObligationsStatus = "pending"
	PARTIALLY_SETTLED FinancialObligationsStatus = "partially_settled"
	SETTLED           FinancialObligationsStatus = "settled"
	CANCELED          FinancialObligationsStatus = "canceled"
)

type FinancialObligationsTypes string

const (
	RECEIVABLE FinancialObligationsTypes = "receivable"
	PAYABLE    FinancialObligationsTypes = "payable"
)

type FinancialObligationModel struct {
	Id             int
	CategoryId     int
	Description    string
	Type           FinancialObligationsTypes
	Status         FinancialObligationsStatus
	OriginalAmount float64
	DueDate        time.Time
	CompetenceDate *time.Time
	Notes          *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}
