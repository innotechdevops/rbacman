package group

import (
	"github.com/innotechdevops/rbacman/pkg/core"
)

type Group struct {
	Id       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	ParentId string `json:"parent_id" db:"parent_id"`
}

type CreateGroup struct {
	Id       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	ParentId string `json:"parent_id" db:"parent_id"`
}

type UpdateGroup struct {
	Id       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	ParentId string `json:"parent_id" db:"parent_id"`
}

type DeleteGroup struct {
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
