package organization

import (
	"github.com/innotechdevops/rbacman/pkg/core"
)

type Organization struct {
	ParentId int64  `json:"parent_id" db:"parent_id"`
	Id       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
}

type CreateOrganization struct {
	ParentId int64  `json:"parent_id" db:"parent_id"`
	Id       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
}

type UpdateOrganization struct {
	ParentId int64  `json:"parent_id" db:"parent_id"`
	Id       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
}

type DeleteOrganization struct {
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
