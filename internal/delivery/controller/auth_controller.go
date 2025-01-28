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
