package user_role

import (
	"fmt"
	"github.com/innotechdevops/rbacman/internal/pkg/response"
	"github.com/innotechdevops/rbacman/internal/rbacman/database"
	"github.com/prongbang/sqlxwrapper/mrwrapper"
)

type DataSource interface {
	Count(params Params) int64
	FindList(params Params) []UserRole
	FindById(id int64) UserRole
	Create(data *CreateUserRole) error
	Update(data *UpdateUserRole) error
	Delete(id int64) error
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) Count(params Params) int64 {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT COUNT(id) FROM users_roles WHERE 1=1 %s"
	wheres := ""
	args := []any{}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.Count(conn, sql, args...)
}

func (d *dataSource) FindList(params Params) []UserRole {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT u.id, u.user_id, u.role_id FROM users_roles u WHERE 1=1 %s ORDER BY u.id"
	wheres := ""
	args := []any{}

	if params.Limit > 0 && params.Offset >= 0 {
		sql += " LIMIT ?"
		args = append(args, params.Limit)
		sql += " OFFSET ?"
		args = append(args, params.Offset)
	}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.SelectList[UserRole](conn, sql, args...)
}

func (d *dataSource) FindById(id int64) UserRole {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT u.id, u.user_id, u.role_id FROM users_roles u WHERE u.id = ?"

	return mrwrapper.SelectOne[UserRole](conn, sql, id)
}

func (d *dataSource) Create(data *CreateUserRole) error {
	conn := d.Driver.GetMariaDB()
	sql := "INSERT INTO users_roles (user_id, role_id) VALUES (?, ?)"
	args := []any{
		data.UserId,
		data.RoleId,
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

func (d *dataSource) Update(data *UpdateUserRole) error {
	conn := d.Driver.GetMariaDB()
	params := map[string]interface{}{
		"id": data.Id,
	}
	set := ""
	sql := "UPDATE users_roles SET %s WHERE id=:id"

	if data.UserId != "" {
		set += ", user_id=:user_id"
		params["user_id"] = data.UserId
	}
	if data.RoleId != "" {
		set += ", role_id=:role_id"
		params["role_id"] = data.RoleId
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
	sql := "DELETE FROM users_roles WHERE id=?"

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
