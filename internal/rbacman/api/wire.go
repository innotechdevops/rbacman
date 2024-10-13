//go:build wireinject
// +build wireinject

package api

import (
	"github.com/google/wire"
	"github.com/innotechdevops/rbacman/internal/pkg/response"
	"github.com/innotechdevops/rbacman/internal/rbacman/database"

	"github.com/innotechdevops/rbacman/internal/rbacman/api/permission"
	"github.com/innotechdevops/rbacman/internal/rbacman/api/role"
	//+fibergen:import wire:package
)

func CreateAPI(dbDriver database.Drivers) API {
	wire.Build(
		NewAPI,
		NewRouters,
		response.New,
		role.ProviderSet,
		permission.ProviderSet,
		//+fibergen:func wire:build
	)
	return nil
}
