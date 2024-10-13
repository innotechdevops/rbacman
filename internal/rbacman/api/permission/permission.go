package permission

import (
	"github.com/innotechdevops/rbacman/pkg/core"
)

type Permission struct {
	Id    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Value string `json:"value" db:"value"`
}

type CreatePermission struct {
	Id    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Value string `json:"value" db:"value"`
}

type UpdatePermission struct {
	Id    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Value string `json:"value" db:"value"`
}

type DeletePermission struct {
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