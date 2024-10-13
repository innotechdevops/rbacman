package user_role

import (
	"github.com/innotechdevops/rbacman/pkg/core"
)

type UserRole struct {
	Id     int64  `json:"id" db:"id"`
	UserId string `json:"user_id" db:"user_id"`
	RoleId string `json:"role_id" db:"role_id"`
}

type CreateUserRole struct {
	Id     int64  `json:"id" db:"id"`
	UserId string `json:"user_id" db:"user_id"`
	RoleId string `json:"role_id" db:"role_id"`
}

type UpdateUserRole struct {
	Id     int64  `json:"id" db:"id"`
	UserId string `json:"user_id" db:"user_id"`
	RoleId string `json:"role_id" db:"role_id"`
}

type DeleteUserRole struct {
	Id int64 `json:"id"`
}

type QueryOne struct {
	Id int64 `json:"id"`
}

type QueryMany struct {
	core.Params
}

type Params struct {
	QueryOne
	QueryMany
}
