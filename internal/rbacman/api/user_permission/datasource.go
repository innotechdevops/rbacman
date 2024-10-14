package user_permission

import (
	"github.com/innotechdevops/rbacman/internal/rbacman/database"
	"github.com/prongbang/sqlxwrapper/pqwrapper"
)

type DataSource interface {
	PermissionAllowed(userId string, resourcePermission string) *UserPermission
	PermissionList(userId string) []UserPermission
}

type dataSource struct {
	Driver database.Drivers
}

// PermissionAllowed implements DataSource.
func (r *dataSource) PermissionAllowed(userId string, resourcePermission string) *UserPermission {
	conn := r.Driver.GetMariaDB()
	query := `SELECT 
		g.name AS group_name,
		r.name AS resource_name,
		r.value AS resource_value,
		p.name AS permission_name,
		p.value AS permission_value,
		CONCAT(r.value, '_', p.value) AS resource_permission
	FROM users u 
	INNER JOIN users_groups ug ON ug.users_id = u.id
	INNER JOIN groups g ON g.id = ug.groups_id
	INNER JOIN groups_permissions gp ON gp.groups_id = ug.groups_id
	INNER JOIN permissions p ON p.id = gp.permissions_id
	INNER JOIN resources r ON r.id = gp.resources_id
	WHERE u.id = ? AND CONCAT(r.value, '_', p.value) = ?`

	args := []any{userId, resourcePermission}

	return pqwrapper.SelectOne[*UserPermission](conn, query, args...)
}

// PermissionList implements DataSource.
func (r *dataSource) PermissionList(userId string) []UserPermission {
	conn := r.Driver.GetMariaDB()
	query := `SELECT 
		g.name AS group_name,
		r.name AS resource_name,
		r.value AS resource_value,
		p.name AS permission_name,
		p.value AS permission_value,
		CONCAT(r.value, '_', p.value) AS resource_permission
	FROM users u 
	INNER JOIN users_groups ug ON ug.users_id = u.id
	INNER JOIN groups g ON g.id = ug.groups_id
	INNER JOIN groups_permissions gp ON gp.groups_id = ug.groups_id
	INNER JOIN permissions p ON p.id = gp.permissions_id
	INNER JOIN resources r ON r.id = gp.resources_id
	WHERE u.id = ?`

	args := []any{userId}

	return pqwrapper.SelectList[UserPermission](conn, query, args...)
}

func NewDataSource(driver database.Drivers) DataSource {
	return &dataSource{
		Driver: driver,
	}
}
