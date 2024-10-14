package user

import (
	"fmt"
	"github.com/innotechdevops/rbacman/internal/pkg/response"
	"github.com/innotechdevops/rbacman/internal/rbacman/database"
	"github.com/prongbang/sqlxwrapper/mrwrapper"
)

type DataSource interface {
	Count(params Params) int64
	FindList(params Params) []User
	FindById(id string) User
	Create(data *CreateUser) error
	Update(data *UpdateUser) error
	Delete(id string) error
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) Count(params Params) int64 {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT COUNT(id) FROM users WHERE 1=1 %s"
	wheres := ""
	args := []any{}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.Count(conn, sql, args...)
}

func (d *dataSource) FindList(params Params) []User {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT u.password, u.id, u.first_name, u.last_name, u.username, u.email FROM users u WHERE 1=1 %s ORDER BY u.id"
	wheres := ""
	args := []any{}

	if params.Limit > 0 && params.Offset >= 0 {
		sql += " LIMIT ?"
		args = append(args, params.Limit)
		sql += " OFFSET ?"
		args = append(args, params.Offset)
	}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.SelectList[User](conn, sql, args...)
}

func (d *dataSource) FindById(id string) User {
	conn := d.Driver.GetMariaDB()
	sql := "SELECT u.password, u.id, u.first_name, u.last_name, u.username, u.email FROM users u WHERE u.id = ?"

	return mrwrapper.SelectOne[User](conn, sql, id)
}

func (d *dataSource) Create(data *CreateUser) error {
	conn := d.Driver.GetMariaDB()
	sql := "INSERT INTO users (password, first_name, last_name, username, email) VALUES (?, ?, ?, ?, ?)"
	args := []any{
		data.Password,
		data.FirstName,
		data.LastName,
		data.Username,
		data.Email,
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

func (d *dataSource) Update(data *UpdateUser) error {
	conn := d.Driver.GetMariaDB()
	params := map[string]interface{}{
		"id": data.Id,
	}
	set := ""
	sql := "UPDATE users SET %s WHERE id=:id"

	if data.Password != "" {
		set += ", password=:password"
		params["password"] = data.Password
	}
	if data.FirstName != "" {
		set += ", first_name=:first_name"
		params["first_name"] = data.FirstName
	}
	if data.LastName != "" {
		set += ", last_name=:last_name"
		params["last_name"] = data.LastName
	}
	if data.Username != "" {
		set += ", username=:username"
		params["username"] = data.Username
	}
	if data.Email != "" {
		set += ", email=:email"
		params["email"] = data.Email
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
	sql := "DELETE FROM users WHERE id=?"

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
