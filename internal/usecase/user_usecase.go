package usecase

import (
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/entity"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/exception"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/lib/password"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/lib/storage"
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

	return model.UserResource(ctx, user), nil
}

func (u *UserUseCase) GetAllUser(ctx *fiber.Ctx, request *model.GetAllUserRequest) (*model.PaginationResponse, error) {
	tx := u.DB.Begin()
	defer tx.Rollback()

	user := new([]entity.User)

	users, meta, err := u.UserRepository.Paginate(tx, ctx, user, request.Page, request.PageSize)
	if err != nil {
		return nil, exception.Error(fiber.ErrInternalServerError, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Error(fiber.ErrInternalServerError, err.Error())
	}

	responses := make([]model.UserResponse, len(users))
	for i, user := range users {
		responses[i] = *model.UserResource(ctx, &user)
	}

	return &model.PaginationResponse{Data: responses, Meta: meta}, nil
}

func (u *UserUseCase) UpdateUser(ctx *fiber.Ctx, request *model.UpdateUserRequest) (*model.UserResponse, error) {
	tx := u.DB.Begin()
	defer tx.Rollback()

	if err := u.Validate.Struct(request); err != nil {
		return nil, exception.Validate(fiber.ErrUnprocessableEntity, err)
	}

	user := new(entity.User)
	if err := u.UserRepository.FindById(tx, ctx, user, request.ID); err != nil {
		return nil, exception.Error(fiber.ErrNotFound, err.Error())
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Password != "" {
		hashedPassword, err := password.Hash(request.Password)
		if err != nil {
			return nil, exception.Error(fiber.ErrInternalServerError, err.Error())
		}

		user.Password = hashedPassword
	}

	file, err := ctx.FormFile("photo")
	if err == nil {
		mimeType := []string{"image/jpeg", "image/png", "image/webp"}
		filePath, err := storage.UploadSingle(ctx, storage.Public, file, mimeType, "user")
		if err != nil {
			return nil, exception.Error(fiber.ErrInternalServerError, err.Error())
		}

		user.Photo = filePath
	}

	if err := u.UserRepository.Update(tx, ctx, user, request.ID); err != nil {
		return nil, exception.Error(fiber.ErrInternalServerError, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Error(fiber.ErrInternalServerError, err.Error())
	}

	return model.UserResource(ctx, user), nil
}

func (u *UserUseCase) DeleteUserById(ctx *fiber.Ctx, id uint) (*model.UserResponse, *exception.ErrorResponse) {
	tx := u.DB.Begin()
	defer tx.Rollback()

	user := new(entity.User)
	if err := u.UserRepository.FindById(tx, ctx, user, id); err != nil {
		return nil, exception.Error(fiber.ErrNotFound, err.Error())
	}

	if err := u.UserRepository.Delete(tx, ctx, user, id); err != nil {
		return nil, exception.Error(fiber.ErrInternalServerError, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Error(fiber.ErrInternalServerError, err.Error())
	}

	return model.UserResource(ctx, user), nil
}
