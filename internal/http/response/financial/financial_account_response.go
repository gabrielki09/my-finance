package financialresponse

import (
	financialmodel "finance/models/financial"
	"time"
)

type FinancialAccountResponse struct {
	Id             int                                 `json:"id"`
	Name           string                              `json:"name"`
	Type           financialmodel.FinancialAccountType `json:"type"`
	InitialBalance float64                             `json:"initial_balance"`
	OpenedAt       time.Time                           `json:"opened_at"`
	CreatedAt      time.Time                           `json:"created_at"`
	UpdatedAt      time.Time                           `json:"updated_at"`
}
