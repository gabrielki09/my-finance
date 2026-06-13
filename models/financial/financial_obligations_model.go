package financialmodel

import "time"

type FinancialObligationsStatus string

const (
	PENDING           FinancialObligationsStatus = "pending"
	PARTIALLY_SETTLED FinancialObligationsStatus = "partially_settled"
	SETTLED           FinancialObligationsStatus = "settled"
	CANCELED          FinancialObligationsStatus = "canceled"
)

type FinancialObligationsTpyes string

const (
	PAYABLE    FinancialObligationsTpyes = "payable"
	RECEIVABLE FinancialObligationsTpyes = "receivable"
)

type FinancialObligationModel struct {
	Id          int
	CategoryId  string
	Description string
	Type        FinancialObligationsTpyes
	Status      FinancialObligationsStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
