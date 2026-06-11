package tenant_user_model

import "time"

type TenantModel struct {
	Id        string     `json:"id"`
	TenantId  string     `json:"tenant_id"`
	UserId    string     `json:"user_id"`
	Role      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
