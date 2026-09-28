package categorymodel

import "time"

type CategoryTpyes string

const (
	INCOME  CategoryTpyes = "income"
	EXPENSE CategoryTpyes = "expense"
	BOTH    CategoryTpyes = "both"
)

type CategoryModel struct {
	Id        int
	ParentId  *int
	Name      string
	Type      CategoryTpyes
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
