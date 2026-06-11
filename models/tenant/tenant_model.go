package tenant_model

import "time"

type TenantModel struct {
	Id        string
	Domain    *string
	Document  *string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
