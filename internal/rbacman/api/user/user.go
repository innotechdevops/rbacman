package user

import (
	"github.com/innotechdevops/rbacman/pkg/core"
)

type User struct {
	Password  string `json:"password" db:"password"`
	Id        string `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Username  string `json:"username" db:"username"`
	Email     string `json:"email" db:"email"`
}

type CreateUser struct {
	Password  string `json:"password" db:"password"`
	Id        string `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Username  string `json:"username" db:"username"`
	Email     string `json:"email" db:"email"`
}

type UpdateUser struct {
	Password  string `json:"password" db:"password"`
	Id        string `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Username  string `json:"username" db:"username"`
	Email     string `json:"email" db:"email"`
}

type DeleteUser struct {
	Id string `json:"id"`
}

type QueryOne struct {
	Id string `json:"id"`
}

type QueryMany struct {
	core.Params
}

type Params struct {
	QueryOne
	QueryMany
}
