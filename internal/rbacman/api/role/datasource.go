package role

import (
	"fmt"
	"github.com/innotechdevops/rbacman/internal/pkg/response"
	"github.com/innotechdevops/rbacman/internal/rbacman/database"
	"github.com/prongbang/sqlxwrapper/mrwrapper"
)

type DataSource interface {
	Count(params Params) int64
	FindList(params Params) []Role
	FindById(id int64) Role
	Create(data *CreateRole) error
	Update(data *UpdateRole) error
	Delete(id int64) error
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) Count(params Params) int64 {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT COUNT(id) FROM roles WHERE 1=1 %s"
	wheres := ""
	args := []any{}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.Count(conn, sql, args...)
}

func (d *dataSource) FindList(params Params) []Role {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT r.value, r.id, r.name FROM roles r WHERE 1=1 %s ORDER BY r.id"
	wheres := ""
	args := []any{}

	if params.Limit > 0 && params.Offset >= 0 {
		sql += " LIMIT ?"
		args = append(args, params.Limit)
		sql += " OFFSET ?"
		args = append(args, params.Offset)
	}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.SelectList[Role](conn, sql, args...)
}

func (d *dataSource) FindById(id int64) Role {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT r.value, r.id, r.name FROM roles r WHERE r.id = ?"

	return mrwrapper.SelectOne[Role](conn, sql, id)
}

func (d *dataSource) Create(data *CreateRole) error {
	conn := d.Driver.GetMariaDB()
	sql := "INSERT INTO roles (value, name) VALUES (?, ?)"
	args := []any{
		data.Value,
		data.Name,
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

func (d *dataSource) Update(data *UpdateRole) error {
	conn := d.Driver.GetMariaDB()
	params := map[string]interface{}{
		"id": data.Id,
	}
	set := ""
	sql := "UPDATE roles SET %s WHERE id=:id"

	if data.Value != "" {
		set += ", value=:value"
		params["value"] = data.Value
	}
	if data.Name != "" {
		set += ", name=:name"
		params["name"] = data.Name
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
	sql := "DELETE FROM roles WHERE id=?"

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
