package controller

import (
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/exception"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/model"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase) *UserController {
	return &UserController{
		UseCase: useCase,
	}
}

// @Summary      Get all user
// @Description  Get list of user
// @Tags         User
// @Accept       json
// @Produce      json
// @Security Bearer
// @Param page query int false "Page"
// @Param page_size query int false "Page Size"
// @Success 200 {array} model.PaginationResponse{}
// @Failure 500 {object} model.Response{}
// @Router       /api/users [get]
func (c *UserController) GetAllUser(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	pageSize := ctx.QueryInt("page_size", 10)

	request := new(model.GetAllUserRequest)
	request.Page = page
	request.PageSize = pageSize

	response, err := c.UseCase.GetAllUser(ctx, request)
	if err != nil {
		return err
	}

	return model.PageResponse(ctx, model.StatusOK, response)
}

// @Summary      Get user
// @Description  Get user by user id
// @Tags         User
// @Accept       json
// @Produce      json
// @Security Bearer
// @Param id path int true "User ID"
// @Success 200 {object} model.UserResponse{}
// @Failure 400 {object} model.Response{}
// @Failure 401 {object} model.Response{}
// @Failure 422 {object} model.Response{}
// @Failure 404 {object} model.Response{}
// @Failure 500 {object} model.Response{}
// @Router       /api/users/{id} [get]
func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return exception.Error(fiber.ErrBadRequest, "The id param must be of integer type")
	}

	request := &model.GetUserRequest{
		ID: uint(id),
	}

	response, err := c.UseCase.GetUserById(ctx, request)
	if err != nil {
		return err
	}

	return model.WebResponse(ctx, model.StatusOK, response)
}

// @Summary      Update user
// @Description  Update user by user id
// @Tags         User
// @Accept       json
// @Produce      json
// @Security Bearer
// @Param id path int true "User ID"
// @Param name formData string false "Name"
// @Param password formData string false "Password"
// @Param photo formData file false "Photo"
// @Success 200 {object} model.UserResponse{}
// @Failure 400 {object} model.Response{}
// @Failure 401 {object} model.Response{}
// @Failure 422 {object} model.Response{}
// @Failure 404 {object} model.Response{}
// @Failure 500 {object} model.Response{}
// @Router       /api/users/{id} [patch]
func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id, e := ctx.ParamsInt("id")
	if e != nil {
		return exception.Error(fiber.ErrBadRequest, "The id param must be of integer type")
	}

	request := new(model.UpdateUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		return exception.Error(fiber.ErrBadRequest, err.Error())
	}

	request.ID = uint(id)
	response, err := c.UseCase.UpdateUser(ctx, request)
	if err != nil {
		return err
	}

	return model.WebResponse(ctx, model.StatusOK, response)
}

// @Summary      Delete user
// @Description  Delete user account by user id
// @Tags         User
// @Accept       json
// @Produce      json
// @Security Bearer
// @Param id path int true "User ID"
// @Success 204 {object} model.UserResponse{}
// @Failure 400 {object} model.Response{}
// @Failure 401 {object} model.Response{}
// @Failure 404 {object} model.Response{}
// @Failure 500 {object} model.Response{}
// @Router       /api/users/{id} [delete]
func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id, e := ctx.ParamsInt("id")
	if e != nil {
		return exception.Error(fiber.ErrBadRequest, "The id param must be of integer type")
	}

	response, err := c.UseCase.DeleteUserById(ctx, uint(id))
	if err != nil {
		return err
	}

	return model.WebResponse(ctx, model.StatusNoContent, response)
}
