package group_permission

import (
	"github.com/gofiber/fiber/v2"
	"github.com/innotechdevops/rbacman/pkg/core"
)

type Router interface {
	core.Router
}

type router struct {
	Handle   Handler
	Validate Validate
}

func (r *router) Initial(app *fiber.App) {
	v1 := app.Group("/v1")
	{
		v1.Post("/group-permission/one", r.Validate.FindById, r.Handle.FindById)
		v1.Post("/group-permission/many", r.Validate.FindList, r.Handle.FindList)
		v1.Post("/group-permission/create", r.Validate.Create, r.Handle.Create)
		v1.Post("/group-permission/update", r.Validate.Update, r.Handle.Update)
		v1.Post("/group-permission/delete", r.Validate.Delete, r.Handle.Delete)
	}
}

func NewRouter(handle Handler, validate Validate) Router {
	return &router{
		Handle:   handle,
		Validate: validate,
	}
}
