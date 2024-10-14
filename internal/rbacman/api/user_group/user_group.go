package user_group

import (
	"github.com/innotechdevops/rbacman/pkg/core"
)

type UserGroup struct {
	Id      int64  `json:"id" db:"id"`
	UserId  string `json:"userId" db:"user_id"`
	GroupId string `json:"groupId" db:"group_id"`
}

type CreateUserGroup struct {
	Id      int64  `json:"id" db:"id"`
	UserId  string `json:"userId" db:"user_id"`
	GroupId string `json:"groupId" db:"group_id"`
}

type UpdateUserGroup struct {
	Id      int64  `json:"id" db:"id"`
	UserId  string `json:"userId" db:"user_id"`
	GroupId string `json:"groupId" db:"group_id"`
}

type DeleteUserGroup struct {
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
