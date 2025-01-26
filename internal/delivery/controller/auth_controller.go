package controller

import "github.com/dickidarmawansaputra/go-clean-architecture/internal/usecase"

type AuthController struct {
	UseCase *usecase.AuthUseCase
}

func NewAuthController(useCase *usecase.AuthUseCase) *AuthController {
	return &AuthController{
		UseCase: useCase,
	}
}
