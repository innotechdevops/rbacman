package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/innotechdevops/rbacman/internal/rbacman/api/permission"
	"github.com/innotechdevops/rbacman/internal/rbacman/api/role"
	"github.com/innotechdevops/rbacman/pkg/core"
	//+fibergen:import routers:package
)

type Routers interface {
	core.Routers
}

type routers struct {
	RoleRoute       role.Router
	PermissionRoute permission.Router
	//+fibergen:struct routers
}

func (r *routers) Initials(app *fiber.App) {
	r.RoleRoute.Initial(app)
	r.PermissionRoute.Initial(app)
	//+fibergen:func initials
}

func NewRouters(
	roleRoute role.Router,
	permissionRoute permission.Router,
	//+fibergen:func new:routers
) Routers {
	return &routers{
		RoleRoute:       roleRoute,
		PermissionRoute: permissionRoute,
		//+fibergen:return &routers
	}
}
