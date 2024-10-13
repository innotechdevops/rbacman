package permission

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
// @Tags permission
// @Summary Find a permission by id
// @Accept json
// @Produce json
// @Param query body QueryOne true "query"
// @Success 200 {object} core.Success{data=Permission}
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/permission/one [post]
func (h *handler) FindById(c *fiber.Ctx) error {
	q := QueryOne{}
	_ = c.BodyParser(&q)

	if r := h.Uc.FindById(q.Id); r.Id > 0 {
		return h.Response.With(c).Response(goerror.NewOK(r))
	}

	return h.Response.With(c).Response(goerror.NewNotFound())
}

// FindList
// @Tags permission
// @Summary Find a list of permission
// @Accept json
// @Produce json
// @Param query body QueryMany true "query"
// @Success 200 {object} core.Success{data=core.Paging{list=[]Permission}}
// @Failure 400 {object} core.Error
// @Security JWTAuth
// @Router /v1/permission/many [post]
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
// @Tags permission
// @Summary Create a permission
// @Accept json
// @Produce json
// @Param permission body CreatePermission true "permission"
// @Success 201 {object} core.Success{data=Permission}
// @Failure 400 {object} core.Error
// @Security JWTAuth
// @Router /v1/permission/create [post]
func (h *handler) Create(c *fiber.Ctx) error {
	b := CreatePermission{}
	_ = c.BodyParser(&b)

	d, err := h.Uc.Create(&b)
	if err != nil {
		return h.Response.With(c).Response(err)
	}

	return h.Response.With(c).Response(goerror.NewCreated(d))
}

// Update
// @Tags permission
// @Summary Update a permission
// @Accept json
// @Produce json
// @Param permission body UpdatePermission true "permission"
// @Success 200 {object} core.Success{data=Permission}
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/permission/update [post]
func (h *handler) Update(c *fiber.Ctx) error {
	b := UpdatePermission{}
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
// @Tags permission
// @Summary Delete a permission by id
// @Accept json
// @Produce json
// @Param permission body DeletePermission true "permission"
// @Success 200 {object} core.Success
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/permission/delete [post]
func (h *handler) Delete(c *fiber.Ctx) error {
	b := DeletePermission{}
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
