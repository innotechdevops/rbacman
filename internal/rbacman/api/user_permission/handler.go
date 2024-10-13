package user_permission

import (
	"github.com/gofiber/fiber/v2"
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
// @Tags user_permission
// @Summary Find a user_permission by id
// @Accept json
// @Produce json
// @Param query body QueryOne true "query"
// @Success 200 {object} core.Success{data=UserPermission}
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/user-permission/one [post]
func (h *handler) FindById(c *fiber.Ctx) error {
	return h.Response.With(c).Response(goerror.NewOK(nil))
}

// FindList
// @Tags user_permission
// @Summary Find a list of user_permission
// @Accept json
// @Produce json
// @Param query body QueryMany true "query"
// @Success 200 {object} core.Success{data=core.Paging{list=[]UserPermission}}
// @Failure 400 {object} core.Error
// @Security JWTAuth
// @Router /v1/user-permission/many [post]
func (h *handler) FindList(c *fiber.Ctx) error {
	return h.Response.With(c).Response(goerror.NewOK(nil))
}

// Create
// @Tags user_permission
// @Summary Create a user_permission
// @Accept json
// @Produce json
// @Param user_permission body CreateUserPermission true "user_permission"
// @Success 201 {object} core.Success{data=UserPermission}
// @Failure 400 {object} core.Error
// @Security JWTAuth
// @Router /v1/user-permission/create [post]
func (h *handler) Create(c *fiber.Ctx) error {
	return h.Response.With(c).Response(goerror.NewOK(nil))
}

// Update
// @Tags user_permission
// @Summary Update a user_permission
// @Accept json
// @Produce json
// @Param user_permission body UpdateUserPermission true "user_permission"
// @Success 200 {object} core.Success{data=UserPermission}
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/user-permission/update [post]
func (h *handler) Update(c *fiber.Ctx) error {
	return h.Response.With(c).Response(goerror.NewOK(nil))
}

// Delete
// @Tags user_permission
// @Summary Delete a user_permission by id
// @Accept json
// @Produce json
// @Param user_permission body DeleteUserPermission true "user_permission"
// @Success 200 {object} core.Success
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security JWTAuth
// @Router /v1/user-permission/delete [post]
func (h *handler) Delete(c *fiber.Ctx) error {
	return h.Response.With(c).Response(goerror.NewOK(nil))
}

func NewHandler(uc UseCase, response fibererror.Response) Handler {
	return &handler{
		Uc:       uc,
		Response: response,
	}
}
