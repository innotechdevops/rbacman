package resource

import (
	"github.com/innotechdevops/rbacman/pkg/core"
)

type Resource struct {
	Value string `json:"value" db:"value"`
	Id    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
}

type CreateResource struct {
	Value string `json:"value" db:"value"`
	Id    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
}

type UpdateResource struct {
	Value string `json:"value" db:"value"`
	Id    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
}

type DeleteResource struct {
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
