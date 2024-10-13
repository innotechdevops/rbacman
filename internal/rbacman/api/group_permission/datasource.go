package group_permission

import (
	"fmt"
	"github.com/innotechdevops/rbacman/internal/pkg/response"
	"github.com/innotechdevops/rbacman/internal/rbacman/database"
	"github.com/prongbang/sqlxwrapper/mrwrapper"
)

type DataSource interface {
	Count(params Params) int64
	FindList(params Params) []GroupPermission
	FindById(id int64) GroupPermission
	Create(data *CreateGroupPermission) error
	Update(data *UpdateGroupPermission) error
	Delete(id int64) error
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) Count(params Params) int64 {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT COUNT(id) FROM group_permission WHERE 1=1 %s"
	wheres := ""
	args := []any{}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.Count(conn, sql, args...)
}

func (d *dataSource) FindList(params Params) []GroupPermission {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT g.id, g.group_id, g.resource_id, g.permission_id FROM group_permission g WHERE 1=1 %s ORDER BY g.id"
	wheres := ""
	args := []any{}

	if params.Limit > 0 && params.Offset >= 0 {
		sql += " LIMIT ?"
		args = append(args, params.Limit)
		sql += " OFFSET ?"
		args = append(args, params.Offset)
	}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.SelectList[GroupPermission](conn, sql, args...)
}

func (d *dataSource) FindById(id int64) GroupPermission {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT g.id, g.group_id, g.resource_id, g.permission_id FROM group_permission g WHERE g.id = ?"

	return mrwrapper.SelectOne[GroupPermission](conn, sql, id)
}

func (d *dataSource) Create(data *CreateGroupPermission) error {
	conn := d.Driver.GetMariaDB()
	sql := "INSERT INTO group_permission (group_id, resource_id, permission_id) VALUES (?, ?, ?)"
	args := []any{
		data.GroupId,
		data.ResourceId,
		data.PermissionId,
	}
	tx, err := mrwrapper.Create(conn, sql, []any{&data.Id}, args...)
	if err == nil {
		if e := tx.Commit(); e != nil {
			return response.NewCommitError()
		}
		return nil
	}
	return response.NewInsertError()
}

func (d *dataSource) Update(data *UpdateGroupPermission) error {
	conn := d.Driver.GetMariaDB()
	params := map[string]interface{}{
		"id": data.Id,
	}
	set := ""
	sql := "UPDATE group_permission SET %s WHERE id=:id"

	if data.GroupId != "" {
		set += ", group_id=:group_id"
		params["group_id"] = data.GroupId
	}
	if data.ResourceId > 0 {
		set += ", resource_id=:resource_id"
		params["resource_id"] = data.ResourceId
	}
	if data.PermissionId > 0 {
		set += ", permission_id=:permission_id"
		params["permission_id"] = data.PermissionId
	}

	tx, err := mrwrapper.Update(conn, sql, set, params)
	if err == nil {
		if e := tx.Commit(); e != nil {
			return response.NewCommitError()
		}
		return nil
	}
	return response.NewUpdateError()
}

func (d *dataSource) Delete(id int64) error {
	conn := d.Driver.GetMariaDB()
	sql := "DELETE FROM group_permission WHERE id=?"

	tx, err := mrwrapper.Delete(conn, sql, id)
	if err == nil {
		if e := tx.Commit(); e != nil {
			return response.NewCommitError()
		}
		return nil
	}
	return response.NewDeleteError()
}

func NewDataSource(driver database.Drivers) DataSource {
	return &dataSource{
		Driver: driver,
	}
}
