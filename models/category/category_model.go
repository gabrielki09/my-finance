package categorymodel

import (
	"time"
)

type CategoryModel struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
