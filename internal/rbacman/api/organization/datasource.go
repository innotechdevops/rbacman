package organization

import (
	"fmt"
	"github.com/innotechdevops/rbacman/internal/pkg/response"
	"github.com/innotechdevops/rbacman/internal/rbacman/database"
	"github.com/prongbang/sqlxwrapper/mrwrapper"
)

type DataSource interface {
	Count(params Params) int64
	FindList(params Params) []Organization
	FindById(id int64) Organization
	Create(data *CreateOrganization) error
	Update(data *UpdateOrganization) error
	Delete(id int64) error
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) Count(params Params) int64 {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT COUNT(id) FROM organizations WHERE 1=1 %s"
	wheres := ""
	args := []any{}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.Count(conn, sql, args...)
}

func (d *dataSource) FindList(params Params) []Organization {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT o.parent_id, o.id, o.name FROM organizations o WHERE 1=1 %s ORDER BY o.id"
	wheres := ""
	args := []any{}

	if params.Limit > 0 && params.Offset >= 0 {
		sql += " LIMIT ?"
		args = append(args, params.Limit)
		sql += " OFFSET ?"
		args = append(args, params.Offset)
	}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.SelectList[Organization](conn, sql, args...)
}

func (d *dataSource) FindById(id int64) Organization {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT o.parent_id, o.id, o.name FROM organizations o WHERE o.id = ?"

	return mrwrapper.SelectOne[Organization](conn, sql, id)
}

func (d *dataSource) Create(data *CreateOrganization) error {
	conn := d.Driver.GetMariaDB()
	sql := "INSERT INTO organizations (parent_id, name) VALUES (?, ?)"
	args := []any{
		data.ParentId,
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

func (d *dataSource) Update(data *UpdateOrganization) error {
	conn := d.Driver.GetMariaDB()
	params := map[string]interface{}{
		"id": data.Id,
	}
	set := ""
	sql := "UPDATE organizations SET %s WHERE id=:id"

	if data.ParentId > 0 {
		set += ", parent_id=:parent_id"
		params["parent_id"] = data.ParentId
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
	sql := "DELETE FROM organizations WHERE id=?"

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
