package financialresponse

import (
	financialmodel "finance/models/financial"
	"time"
)

type FinancialObligationResponse struct {
	Id          int
	CategoryId  string
	Description string
	Type        financialmodel.FinancialObligationsTpyes
	Status      financialmodel.FinancialObligationsStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
