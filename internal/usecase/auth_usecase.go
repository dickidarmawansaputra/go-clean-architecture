package usecase

import (
	"fmt"

	"github.com/dickidarmawansaputra/go-clean-architecture/internal/entity"
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

func (u *AuthUseCase) Register(ctx *fiber.Ctx, request *model.RegisterRequest) (*entity.User, error) {
	tx := u.DB.Begin()
	defer tx.Rollback()

	if err := u.Validate.Struct(request); err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Photo:    fmt.Sprintf("https://ui-avatars.com/api/?name=%s", request.Name),
	}

	userExists := u.UserRepository.CheckUserExists(tx, ctx, user, request.Email)
	if userExists {
		return nil, fiber.NewError(fiber.StatusConflict, "User already exist")
	}

	hashedPassword, err := password.Hash(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	if err := u.UserRepository.Create(tx, ctx, user); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return user, nil
}
