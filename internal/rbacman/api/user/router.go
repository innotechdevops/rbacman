package user

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
		v1.Post("/user/one", r.Validate.FindById, r.Handle.FindById)
		v1.Post("/user/many", r.Validate.FindList, r.Handle.FindList)
		v1.Post("/user/create", r.Validate.Create, r.Handle.Create)
		v1.Post("/user/update", r.Validate.Update, r.Handle.Update)
		v1.Post("/user/delete", r.Validate.Delete, r.Handle.Delete)
	}
}

func NewRouter(handle Handler, validate Validate) Router {
	return &router{
		Handle:   handle,
		Validate: validate,
	}
}