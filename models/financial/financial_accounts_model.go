package financialmodel

import "time"

type FinancialAccountType string

const (
	CASH     FinancialAccountType = "cash"
	CHECKING FinancialAccountType = "checking"
	SAVINGS  FinancialAccountType = "savings"
	DIGITAL  FinancialAccountType = "digital"
)

type FinancialAccountModel struct {
	Id             int
	Name           string
	Type           FinancialAccountType
	InitialBalance float64
	OpenedAt       time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}
