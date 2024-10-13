package api

import (
	"github.com/gofiber/fiber/v2"
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
	"github.com/innotechdevops/rbacman/pkg/core"
	//+fibergen:import routers:package
)

type Routers interface {
	core.Routers
}

type routers struct {
	RoleRoute             role.Router
	PermissionRoute       permission.Router
	GroupRoute            group.Router
	GroupPermissionRoute  group_permission.Router
	OrganizationRoute     organization.Router
	ResourceRoute         resource.Router
	UserRoute             user.Router
	UserGroupRoute        user_group.Router
	UserOrganizationRoute user_organization.Router
	UserRoleRoute         user_role.Router
	UserPermissionRoute   user_permission.Router
	//+fibergen:struct routers
}

func (r *routers) Initials(app *fiber.App) {
	r.RoleRoute.Initial(app)
	r.PermissionRoute.Initial(app)
	r.GroupRoute.Initial(app)
	r.GroupPermissionRoute.Initial(app)
	r.OrganizationRoute.Initial(app)
	r.ResourceRoute.Initial(app)
	r.UserRoute.Initial(app)
	r.UserGroupRoute.Initial(app)
	r.UserOrganizationRoute.Initial(app)
	r.UserRoleRoute.Initial(app)
	r.UserPermissionRoute.Initial(app)
	//+fibergen:func initials
}

func NewRouters(
	roleRoute role.Router,
	permissionRoute permission.Router,
	groupRoute group.Router,
	groupPermissionRoute group_permission.Router,
	organizationRoute organization.Router,
	resourceRoute resource.Router,
	userRoute user.Router,
	userGroupRoute user_group.Router,
	userOrganizationRoute user_organization.Router,
	userRoleRoute user_role.Router,
	userPermissionRoute user_permission.Router,
	//+fibergen:func new:routers
) Routers {
	return &routers{
		RoleRoute:             roleRoute,
		PermissionRoute:       permissionRoute,
		GroupRoute:            groupRoute,
		GroupPermissionRoute:  groupPermissionRoute,
		OrganizationRoute:     organizationRoute,
		ResourceRoute:         resourceRoute,
		UserRoute:             userRoute,
		UserGroupRoute:        userGroupRoute,
		UserOrganizationRoute: userOrganizationRoute,
		UserRoleRoute:         userRoleRoute,
		UserPermissionRoute:   userPermissionRoute,
		//+fibergen:return &routers
	}
}
