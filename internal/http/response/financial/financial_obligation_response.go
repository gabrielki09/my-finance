package financialresponse

import (
	financialmodel "finance/models/financial"
	"time"
)

type FinancialObligationResponse struct {
	Id             int                                       `json:"id"`
	CategoryId     int                                       `json:"category_id"`
	Description    string                                    `json:"description"`
	Type           financialmodel.FinancialObligationsTypes  `json:"type"`
	Status         financialmodel.FinancialObligationsStatus `json:"status"`
	OriginalAmount float64                                   `json:"original_amount"`
	DueDate        time.Time                                 `json:"due_date"`
	CompetenceDate *time.Time                                `json:"competence_date,omitempty"`
	Notes          *string                                   `json:"notes,omitempty"`
	CreatedAt      time.Time                                 `json:"created_at"`
	UpdatedAt      time.Time                                 `json:"updated_at"`
	DeletedAt      *time.Time                                `json:"deleted_at,omitempty"`
}
