package user_permission

import (
	"github.com/innotechdevops/rbacman/internal/rbacman/database"
	"github.com/prongbang/sqlxwrapper/pqwrapper"
	"strings"
)

type DataSource interface {
	PermissionAllowed(userId string, resourcePermission string) bool
	PermissionList(userId string) []UserPermission
}

type dataSource struct {
	Driver database.Drivers
}

// PermissionAllowed implements DataSource.
func (r *dataSource) PermissionAllowed(userId string, resourcePermission string) bool {
	conn := r.Driver.GetMariaDB()

	// High level
	query1 := `SELECT 
	COUNT(g.id)
FROM users u 
INNER JOIN users_groups ug ON ug.user_id = u.id
INNER JOIN groups g ON g.id = ug.group_id
INNER JOIN roles r ON r.id = g.parent_id
WHERE u.id = ?`
	hlAllowed := pqwrapper.Count(conn, query1, userId) > 0
	if hlAllowed {
		return true
	}

	// User level
	query2 := `SELECT 
		COUNT(p.id)
	FROM users u 
	INNER JOIN users_groups ug ON ug.user_id = u.id
	INNER JOIN groups g ON g.id = ug.group_id
	INNER JOIN groups_permissions gp ON gp.group_id = ug.group_id
	INNER JOIN permissions p ON p.id = gp.permission_id
	INNER JOIN resources r ON r.id = gp.resource_id
	WHERE u.id = ? AND UPPER(CONCAT(r.value, ':', p.value)) = ?`
	userAllowed := pqwrapper.Count(conn, query2, userId, strings.ToUpper(resourcePermission)) > 0

	return userAllowed
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
		CONCAT(r.value, ':', p.value) AS resource_permission
	FROM users u 
	INNER JOIN users_groups ug ON ug.user_id = u.id
	INNER JOIN groups g ON g.id = ug.group_id
	INNER JOIN groups_permissions gp ON gp.group_id = ug.group_id
	INNER JOIN permissions p ON p.id = gp.permission_id
	INNER JOIN resources r ON r.id = gp.resource_id
	WHERE u.id = ?`

	args := []any{userId}

	return pqwrapper.SelectList[UserPermission](conn, query, args...)
}

func NewDataSource(driver database.Drivers) DataSource {
	return &dataSource{
		Driver: driver,
	}
}
