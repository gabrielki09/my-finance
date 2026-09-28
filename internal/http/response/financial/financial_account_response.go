package financialresponse

import (
	financialmodel "finance/models/financial"
	"time"
)

type FinancialAccountResponse struct {
	Id             int                                 `json:"id,omitempty"`
	Name           string                              `json:"name,omitempty"`
	Type           financialmodel.FinancialAccountType `json:"type,omitempty"`
	InitialBalance float64                             `json:"initial_balance,omitempty"`
	OpenedAt       time.Time                           `json:"opened_at,omitempty"`
	CreatedAt      time.Time                           `json:"created_at,omitempty"`
	UpdatedAt      time.Time                           `json:"updated_at,omitempty"`
}
