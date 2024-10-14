package user_role

import (
	"github.com/innotechdevops/rbacman/pkg/core"
)

type UserRole struct {
	Id     int64  `json:"id" db:"id"`
	UserId string `json:"userId" db:"user_id"`
	RoleId string `json:"roleId" db:"role_id"`
}

type CreateUserRole struct {
	Id     int64  `json:"id" db:"id"`
	UserId string `json:"userId" db:"user_id"`
	RoleId string `json:"roleId" db:"role_id"`
}

type UpdateUserRole struct {
	Id     int64  `json:"id" db:"id"`
	UserId string `json:"userId" db:"user_id"`
	RoleId string `json:"roleId" db:"role_id"`
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
