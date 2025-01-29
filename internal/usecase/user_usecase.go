package usecase

import (
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/entity"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/exception"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/model"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB             *gorm.DB
	Validate       *validator.Validate
	UserRepository *repository.UserRepository
}

func NewUserUseCase(db *gorm.DB, validate *validator.Validate, userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		DB:             db,
		Validate:       validate,
		UserRepository: userRepository,
	}
}

func (u *UserUseCase) GetUserById(ctx *fiber.Ctx, request *model.GetUserRequest) (*model.UserResponse, error) {
	tx := u.DB.Begin()
	defer tx.Rollback()

	if err := u.Validate.Struct(request); err != nil {
		return nil, exception.Error(fiber.ErrUnprocessableEntity, err)
	}

	user := new(entity.User)
	if err := u.UserRepository.FindById(tx, ctx, user, request.ID); err != nil {
		return nil, exception.Error(fiber.ErrNotFound, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Error(fiber.ErrInternalServerError, err.Error())
	}

	return model.UserResource(user), nil
}
