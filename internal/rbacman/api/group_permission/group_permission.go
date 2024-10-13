package group_permission

import (
	"github.com/innotechdevops/rbacman/pkg/core"
)

type GroupPermission struct {
	Id           int64  `json:"id" db:"id"`
	GroupId      string `json:"group_id" db:"group_id"`
	ResourceId   int64  `json:"resource_id" db:"resource_id"`
	PermissionId int64  `json:"permission_id" db:"permission_id"`
}

type CreateGroupPermission struct {
	Id           int64  `json:"id" db:"id"`
	GroupId      string `json:"group_id" db:"group_id"`
	ResourceId   int64  `json:"resource_id" db:"resource_id"`
	PermissionId int64  `json:"permission_id" db:"permission_id"`
}

type UpdateGroupPermission struct {
	Id           int64  `json:"id" db:"id"`
	GroupId      string `json:"group_id" db:"group_id"`
	ResourceId   int64  `json:"resource_id" db:"resource_id"`
	PermissionId int64  `json:"permission_id" db:"permission_id"`
}

type DeleteGroupPermission struct {
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
