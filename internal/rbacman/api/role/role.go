package role

import (
	"github.com/innotechdevops/rbacman/pkg/core"
)

type Role struct {
	Value string `json:"value" db:"value"`
	Id    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
}

type CreateRole struct {
	Value string `json:"value" db:"value"`
	Id    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
}

type UpdateRole struct {
	Value string `json:"value" db:"value"`
	Id    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
}

type DeleteRole struct {
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
