package user_model

import "time"

type UserModel struct {
	Id        string
	Name      string
	Cpf       string
	Login     string
	Role      string
	Password  string
	BirthDate *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
