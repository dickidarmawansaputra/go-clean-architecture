package usecase

import (
	"fmt"

	"github.com/dickidarmawansaputra/go-clean-architecture/internal/entity"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/exception"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/lib/jwt"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/lib/password"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/model"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AuthUseCase struct {
	Config         *viper.Viper
	DB             *gorm.DB
	Validate       *validator.Validate
	UserRepository *repository.UserRepository
}

func NewAuthUseCase(config *viper.Viper, db *gorm.DB, validate *validator.Validate, repository *repository.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		Config:         config,
		DB:             db,
		Validate:       validate,
		UserRepository: repository,
	}
}

func (u *AuthUseCase) Register(ctx *fiber.Ctx, request *model.RegisterRequest) (*model.UserResponse, error) {
	tx := u.DB.Begin()
	defer tx.Rollback()

	if err := u.Validate.Struct(request); err != nil {
		return nil, exception.Validate(fiber.ErrUnprocessableEntity, err)
	}

	user := &entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Photo:    fmt.Sprintf("https://ui-avatars.com/api/?name=%s", request.Name),
	}

	if userExists := u.UserRepository.CheckUserExists(tx, ctx, user, request.Email); userExists {
		return nil, exception.Error(fiber.ErrConflict, "User already exists")
	}

	hashedPassword, err := password.Hash(user.Password)
	if err != nil {
		return nil, exception.Error(fiber.ErrInternalServerError, err.Error())
	}

	user.Password = hashedPassword

	if err := u.UserRepository.Create(tx, ctx, user); err != nil {
		return nil, exception.Error(fiber.ErrInternalServerError, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Error(fiber.ErrInternalServerError, err.Error())
	}

	return model.UserResource(user), nil
}

func (u *AuthUseCase) Login(ctx *fiber.Ctx, request *model.LoginRequest) (*model.AuthResponse, error) {
	tx := u.DB.Begin()
	defer tx.Rollback()

	if err := u.Validate.Struct(request); err != nil {
		return nil, exception.Validate(fiber.ErrUnprocessableEntity, err)
	}

	user := &entity.User{
		Email:    request.Email,
		Password: request.Password,
	}

	if err := u.UserRepository.FindUserByEmail(tx, ctx, user, request.Email); err != nil {
		return nil, exception.Error(fiber.ErrUnauthorized, "Incorrect user credentials")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Error(fiber.ErrInternalServerError, err.Error())
	}

	if err := password.Verify(user.Password, request.Password); err != nil {
		return nil, exception.Error(fiber.ErrUnauthorized, "Incorrect password")
	}

	token, err := jwt.Generate(u.Config, &jwt.TokenPayload{ID: user.ID})
	if err != nil {
		return nil, exception.Error(fiber.ErrUnauthorized, err.Error())
	}

	return &model.AuthResponse{Token: token}, nil
}
