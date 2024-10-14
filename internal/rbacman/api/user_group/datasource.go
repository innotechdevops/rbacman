package user_group

import (
	"fmt"
	"github.com/innotechdevops/rbacman/internal/pkg/response"
	"github.com/innotechdevops/rbacman/internal/rbacman/database"
	"github.com/prongbang/sqlxwrapper/mrwrapper"
)

type DataSource interface {
	Count(params Params) int64
	FindList(params Params) []UserGroup
	FindById(id int64) UserGroup
	Create(data *CreateUserGroup) error
	Update(data *UpdateUserGroup) error
	Delete(id int64) error
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) Count(params Params) int64 {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT COUNT(id) FROM users_groups WHERE 1=1 %s"
	wheres := ""
	args := []any{}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.Count(conn, sql, args...)
}

func (d *dataSource) FindList(params Params) []UserGroup {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT u.id, u.user_id, u.group_id FROM users_groups u WHERE 1=1 %s ORDER BY u.id"
	wheres := ""
	args := []any{}

	if params.Limit > 0 && params.Offset >= 0 {
		sql += " LIMIT ?"
		args = append(args, params.Limit)
		sql += " OFFSET ?"
		args = append(args, params.Offset)
	}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.SelectList[UserGroup](conn, sql, args...)
}

func (d *dataSource) FindById(id int64) UserGroup {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT u.id, u.user_id, u.group_id FROM users_groups u WHERE u.id = ?"

	return mrwrapper.SelectOne[UserGroup](conn, sql, id)
}

func (d *dataSource) Create(data *CreateUserGroup) error {
	conn := d.Driver.GetMariaDB()
	sql := "INSERT INTO users_groups (user_id, group_id) VALUES (?, ?)"
	args := []any{
		data.UserId,
		data.GroupId,
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

func (d *dataSource) Update(data *UpdateUserGroup) error {
	conn := d.Driver.GetMariaDB()
	params := map[string]interface{}{
		"id": data.Id,
	}
	set := ""
	sql := "UPDATE users_groups SET %s WHERE id=:id"

	if data.UserId != "" {
		set += ", user_id=:user_id"
		params["user_id"] = data.UserId
	}
	if data.GroupId != "" {
		set += ", group_id=:group_id"
		params["group_id"] = data.GroupId
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
	sql := "DELETE FROM users_groups WHERE id=?"

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
