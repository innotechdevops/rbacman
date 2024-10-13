package user_organization

import (
	"fmt"
	"github.com/innotechdevops/rbacman/internal/pkg/response"
	"github.com/innotechdevops/rbacman/internal/rbacman/database"
	"github.com/prongbang/sqlxwrapper/mrwrapper"
)

type DataSource interface {
	Count(params Params) int64
	FindList(params Params) []UserOrganization
	FindById(id int64) UserOrganization
	Create(data *CreateUserOrganization) error
	Update(data *UpdateUserOrganization) error
	Delete(id int64) error
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) Count(params Params) int64 {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT COUNT(id) FROM user_organization WHERE 1=1 %s"
	wheres := ""
	args := []any{}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.Count(conn, sql, args...)
}

func (d *dataSource) FindList(params Params) []UserOrganization {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT u.user_id, u.organization_id, u.id FROM user_organization u WHERE 1=1 %s ORDER BY u.id"
	wheres := ""
	args := []any{}

	if params.Limit > 0 && params.Offset >= 0 {
		sql += " LIMIT ?"
		args = append(args, params.Limit)
		sql += " OFFSET ?"
		args = append(args, params.Offset)
	}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.SelectList[UserOrganization](conn, sql, args...)
}

func (d *dataSource) FindById(id int64) UserOrganization {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT u.user_id, u.organization_id, u.id FROM user_organization u WHERE u.id = ?"

	return mrwrapper.SelectOne[UserOrganization](conn, sql, id)
}

func (d *dataSource) Create(data *CreateUserOrganization) error {
	conn := d.Driver.GetMariaDB()
	sql := "INSERT INTO user_organization (user_id, organization_id) VALUES (?, ?)"
	args := []any{
		data.UserId,
		data.OrganizationId,
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

func (d *dataSource) Update(data *UpdateUserOrganization) error {
	conn := d.Driver.GetMariaDB()
	params := map[string]interface{}{
		"id": data.Id,
	}
	set := ""
	sql := "UPDATE user_organization SET %s WHERE id=:id"

	if data.UserId != "" {
		set += ", user_id=:user_id"
		params["user_id"] = data.UserId
	}
	if data.OrganizationId > 0 {
		set += ", organization_id=:organization_id"
		params["organization_id"] = data.OrganizationId
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
	sql := "DELETE FROM user_organization WHERE id=?"

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
