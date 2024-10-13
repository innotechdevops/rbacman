package role

import (
	"github.com/gofiber/fiber/v2"
	"github.com/innotechdevops/rbacman/pkg/core"
	"github.com/prongbang/fibererror"
	"github.com/prongbang/goerror"
)

type Handler interface {
	FindById(c *fiber.Ctx) error
	FindList(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type handler struct {
	Uc       UseCase
	Response fibererror.Response
}

// FindById
// @Tags role
// @Summary Find a role by id
// @Accept json
// @Produce json
// @Param query body QueryOne true "query"
// @Success 200 {object} core.Success{data=Role}
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/role/one [post]
func (h *handler) FindById(c *fiber.Ctx) error {
	q := QueryOne{}
	_ = c.BodyParser(&q)

	if r := h.Uc.FindById(q.Id); r.Id > 0 {
		return h.Response.With(c).Response(goerror.NewOK(r))
	}

	return h.Response.With(c).Response(goerror.NewNotFound())
}

// FindList
// @Tags role
// @Summary Find a list of role
// @Accept json
// @Produce json
// @Param query body QueryMany true "query"
// @Success 200 {object} core.Success{data=core.Paging{list=[]Role}}
// @Failure 400 {object} core.Error
// @Security JWTAuth
// @Router /v1/role/many [post]
func (h *handler) FindList(c *fiber.Ctx) error {
	q := QueryMany{}
	_ = c.BodyParser(&q)

	params := Params{
		QueryMany: q,
	}

	getCount := func() int64 { return h.Uc.Count(params) }

	getData := func(limit int64, offset int64) interface{} {
		params.Limit = limit
		params.Offset = offset
		return h.Uc.FindList(params)
	}

	r := core.Pagination(q.Page, q.Limit, getCount, getData)

	return h.Response.With(c).Response(goerror.NewOK(r))
}

// Create
// @Tags role
// @Summary Create a role
// @Accept json
// @Produce json
// @Param role body CreateRole true "role"
// @Success 201 {object} core.Success{data=Role}
// @Failure 400 {object} core.Error
// @Security JWTAuth
// @Router /v1/role/create [post]
func (h *handler) Create(c *fiber.Ctx) error {
	b := CreateRole{}
	_ = c.BodyParser(&b)

	d, err := h.Uc.Create(&b)
	if err != nil {
		return h.Response.With(c).Response(err)
	}

	return h.Response.With(c).Response(goerror.NewCreated(d))
}

// Update
// @Tags role
// @Summary Update a role
// @Accept json
// @Produce json
// @Param role body UpdateRole true "role"
// @Success 200 {object} core.Success{data=Role}
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/role/update [post]
func (h *handler) Update(c *fiber.Ctx) error {
	b := UpdateRole{}
	_ = c.BodyParser(&b)

	if r := h.Uc.FindById(b.Id); r.Id == b.Id {
		d, err := h.Uc.Update(&b)
		if err != nil {
			return h.Response.With(c).Response(err)
		}

		return h.Response.With(c).Response(goerror.NewOK(d))
	}

	return h.Response.With(c).Response(goerror.NewNotFound())
}

// Delete
// @Tags role
// @Summary Delete a role by id
// @Accept json
// @Produce json
// @Param role body DeleteRole true "role"
// @Success 200 {object} core.Success
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/role/delete [post]
func (h *handler) Delete(c *fiber.Ctx) error {
	b := DeleteRole{}
	_ = c.BodyParser(&b)

	if r := h.Uc.FindById(b.Id); r.Id == b.Id {
		if err := h.Uc.Delete(b.Id); err != nil {
			return h.Response.With(c).Response(err)
		}

		return h.Response.With(c).Response(goerror.NewOK(nil))
	}

	return h.Response.With(c).Response(goerror.NewNotFound())
}

func NewHandler(uc UseCase, response fibererror.Response) Handler {
	return &handler{
		Uc:       uc,
		Response: response,
	}
}
