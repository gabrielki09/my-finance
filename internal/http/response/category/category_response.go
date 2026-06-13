package categoryresponse

import (
	categorymodel "finance/models/category"
	"time"
)

type CategoryResponse struct {
	Id        int                         `json:"id"`
	ParentId  *int                        `json:"parent_id"`
	Name      string                      `json:"name"`
	Type      categorymodel.CategoryTpyes `json:"type"`
	CreatedAt time.Time                   `json:"created_at"`
	UpdatedAt time.Time                   `json:"updated_at"`
}
