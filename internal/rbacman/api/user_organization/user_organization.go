package user_organization

import (
	"github.com/innotechdevops/rbacman/pkg/core"
)

type UserOrganization struct {
	UserId         string `json:"user_id" db:"user_id"`
	OrganizationId int64  `json:"organization_id" db:"organization_id"`
	Id             int64  `json:"id" db:"id"`
}

type CreateUserOrganization struct {
	UserId         string `json:"user_id" db:"user_id"`
	OrganizationId int64  `json:"organization_id" db:"organization_id"`
	Id             int64  `json:"id" db:"id"`
}

type UpdateUserOrganization struct {
	UserId         string `json:"user_id" db:"user_id"`
	OrganizationId int64  `json:"organization_id" db:"organization_id"`
	Id             int64  `json:"id" db:"id"`
}

type DeleteUserOrganization struct {
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
