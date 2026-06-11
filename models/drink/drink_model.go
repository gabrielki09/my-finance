package drink_model

import "time"

type DrinkModel struct {
	Id             string     `json:"id"`
	TenantId       string     `json:"tenant_id"`
	CategoryId     int        `json:"category_id"`
	Name           string     `json:"name"`
	Value          float64    `json:"value"`
	Amount         float64    `json:"amount"`
	AlcoholContent float64    `json:"alcohol_content"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}
