package useraddress_model

import "time"

type UserAddressModel struct {
	Id           string
	UserId       int
	Neighborhood string
	Street       string
	Number       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
