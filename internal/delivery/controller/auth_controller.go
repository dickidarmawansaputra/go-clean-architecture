package controller

import (
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/exception"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/model"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	UseCase *usecase.AuthUseCase
}

func NewAuthController(useCase *usecase.AuthUseCase) *AuthController {
	return &AuthController{
		UseCase: useCase,
	}
}

// @Summary      Register user
// @Description  Register user to create account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param request body model.RegisterRequest true "Request Body"
// @Success 201 {object} model.UserResponse{}
// @Failure 400 {object} model.Response{}
// @Failure 409 {object} model.Response{}
// @Failure 422 {object} model.Response{}
// @Failure 500 {object} model.Response{}
// @Router       /api/auth/register [post]
func (c *AuthController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterRequest)

	if err := ctx.BodyParser(request); err != nil {
		return exception.Error(fiber.ErrBadRequest, err.Error())
	}

	response, err := c.UseCase.Register(ctx, request)
	if err != nil {
		return err
	}

	return model.WebResponse(ctx, model.StatusCreated, response)
}

// @Summary      Login user
// @Description  Login user to create JWT Token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param request body model.LoginRequest true "Request Body"
// @Success 200 {object} model.AuthResponse{}
// @Failure 400 {object} model.Response{}
// @Failure 422 {object} model.Response{}
// @Failure 401 {object} model.Response{}
// @Failure 500 {object} model.Response{}
// @Router       /api/auth/login [post]
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginRequest)

	if err := ctx.BodyParser(request); err != nil {
		return exception.Error(fiber.ErrBadRequest, err.Error())
	}

	response, err := c.UseCase.Login(ctx, request)
	if err != nil {
		return err
	}

	return model.WebResponse(ctx, model.StatusOK, response)
}
