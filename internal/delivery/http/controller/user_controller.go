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

// @Summary      Get user
// @Description  Get user by id
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
