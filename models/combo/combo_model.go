package combomodel

import "time"

type ComboModel struct {
	Id        string
	TenantId  string
	Name      string
	Value     float64
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ComboItemModel struct {
	Id         string
	TenantId   string
	ComboId    string
	CategoryId *string
	ProductId  *string
	DrinkId    *string
	Value      float64
	Amount     float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
