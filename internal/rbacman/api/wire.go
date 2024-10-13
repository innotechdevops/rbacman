//go:build wireinject
// +build wireinject

package api

import (
	"github.com/google/wire"
	"github.com/innotechdevops/rbacman/internal/pkg/response"
	"github.com/innotechdevops/rbacman/internal/rbacman/database"

	"github.com/innotechdevops/rbacman/internal/rbacman/api/group"
	"github.com/innotechdevops/rbacman/internal/rbacman/api/group_permission"
	"github.com/innotechdevops/rbacman/internal/rbacman/api/organization"
	"github.com/innotechdevops/rbacman/internal/rbacman/api/permission"
	"github.com/innotechdevops/rbacman/internal/rbacman/api/resource"
	"github.com/innotechdevops/rbacman/internal/rbacman/api/role"
	"github.com/innotechdevops/rbacman/internal/rbacman/api/user"
	"github.com/innotechdevops/rbacman/internal/rbacman/api/user_group"
	"github.com/innotechdevops/rbacman/internal/rbacman/api/user_organization"
	"github.com/innotechdevops/rbacman/internal/rbacman/api/user_permission"
	"github.com/innotechdevops/rbacman/internal/rbacman/api/user_role"
	//+fibergen:import wire:package
)

func CreateAPI(dbDriver database.Drivers) API {
	wire.Build(
		NewAPI,
		NewRouters,
		response.New,
		role.ProviderSet,
		permission.ProviderSet,
		group.ProviderSet,
		group_permission.ProviderSet,
		organization.ProviderSet,
		resource.ProviderSet,
		user.ProviderSet,
		user_group.ProviderSet,
		user_organization.ProviderSet,
		user_role.ProviderSet,
		user_permission.ProviderSet,
		//+fibergen:func wire:build
	)
	return nil
}
