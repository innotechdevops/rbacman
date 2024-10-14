package group

import (
	"fmt"
	"github.com/innotechdevops/rbacman/internal/pkg/response"
	"github.com/innotechdevops/rbacman/internal/rbacman/database"
	"github.com/prongbang/sqlxwrapper/mrwrapper"
)

type DataSource interface {
	Count(params Params) int64
	FindList(params Params) []Group
	FindById(id string) Group
	Create(data *CreateGroup) error
	Update(data *UpdateGroup) error
	Delete(id string) error
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) Count(params Params) int64 {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT COUNT(id) FROM groups WHERE 1=1 %s"
	wheres := ""
	args := []any{}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.Count(conn, sql, args...)
}

func (d *dataSource) FindList(params Params) []Group {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT g.id, g.name, g.parent_id FROM groups g WHERE 1=1 %s ORDER BY g.id"
	wheres := ""
	args := []any{}

	if params.Limit > 0 && params.Offset >= 0 {
		sql += " LIMIT ?"
		args = append(args, params.Limit)
		sql += " OFFSET ?"
		args = append(args, params.Offset)
	}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.SelectList[Group](conn, sql, args...)
}

func (d *dataSource) FindById(id string) Group {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT g.id, g.name, g.parent_id FROM groups g WHERE g.id = ?"

	return mrwrapper.SelectOne[Group](conn, sql, id)
}

func (d *dataSource) Create(data *CreateGroup) error {
	conn := d.Driver.GetMariaDB()
	sql := "INSERT INTO groups (name, parent_id) VALUES (?, ?)"
	args := []any{
		data.Name,
		data.ParentId,
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

func (d *dataSource) Update(data *UpdateGroup) error {
	conn := d.Driver.GetMariaDB()
	params := map[string]interface{}{
		"id": data.Id,
	}
	set := ""
	sql := "UPDATE groups SET %s WHERE id=:id"

	if data.Name != "" {
		set += ", name=:name"
		params["name"] = data.Name
	}
	if data.ParentId != "" {
		set += ", parent_id=:parent_id"
		params["parent_id"] = data.ParentId
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

func (d *dataSource) Delete(id string) error {
	conn := d.Driver.GetMariaDB()
	sql := "DELETE FROM groups WHERE id=?"

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
