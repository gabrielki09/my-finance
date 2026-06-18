package financialresponse

import (
	financialmodel "finance/models/financial"
	"time"
)

type FinancialObligationResponse struct {
	Id          int                                       `json:"id"`
	CategoryId  int                                       `json:"category_id"`
	Description string                                    `json:"description"`
	Type        financialmodel.FinancialObligationsTpyes  `json:"type"`
	Status      financialmodel.FinancialObligationsStatus `json:"status"`
	CreatedAt   time.Time                                 `json:"created_at"`
	UpdatedAt   time.Time                                 `json:"updated_at"`
	DeletedAt   *time.Time                                `json:"deleted_at"`
}
