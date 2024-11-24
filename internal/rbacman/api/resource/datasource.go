package resource

import (
	"fmt"
	"github.com/innotechdevops/rbacman/internal/pkg/response"
	"github.com/innotechdevops/rbacman/internal/rbacman/database"
	"github.com/prongbang/sqlxwrapper/mrwrapper"
)

type DataSource interface {
	Count(params Params) int64
	FindList(params Params) []Resource
	FindById(id int64) Resource
	Create(data *CreateResource) error
	Update(data *UpdateResource) error
	Delete(id int64) error
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) Count(params Params) int64 {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT COUNT(id) FROM resources WHERE 1=1 %s"
	wheres := ""
	args := []any{}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.Count(conn, sql, args...)
}

func (d *dataSource) FindList(params Params) []Resource {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT r.code, r.id, r.name FROM resources r WHERE 1=1 %s ORDER BY r.id"
	wheres := ""
	args := []any{}

	if params.Limit > 0 && params.Offset >= 0 {
		sql += " LIMIT ?"
		args = append(args, params.Limit)
		sql += " OFFSET ?"
		args = append(args, params.Offset)
	}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.SelectList[Resource](conn, sql, args...)
}

func (d *dataSource) FindById(id int64) Resource {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT r.code, r.id, r.name FROM resources r WHERE r.id = ?"

	return mrwrapper.SelectOne[Resource](conn, sql, id)
}

func (d *dataSource) Create(data *CreateResource) error {
	conn := d.Driver.GetMariaDB()
	sql := "INSERT INTO resources (code, name) codeS (?, ?)"
	args := []any{
		data.Code,
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

func (d *dataSource) Update(data *UpdateResource) error {
	conn := d.Driver.GetMariaDB()
	params := map[string]interface{}{
		"id": data.Id,
	}
	set := ""
	sql := "UPDATE resources SET %s WHERE id=:id"

	if data.Code != "" {
		set += ", code=:code"
		params["code"] = data.Code
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
	sql := "DELETE FROM resources WHERE id=?"

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
