package model

import (
	"time"

	"github.com/dickidarmawansaputra/go-clean-architecture/internal/entity"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/lib/storage"
	"github.com/gofiber/fiber/v2"
)

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UserResource(ctx *fiber.Ctx, user *entity.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Photo:     storage.Url(ctx, user.Photo),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type GetUserRequest struct {
	ID uint `json:"id" validate:"required"`
}

type GetAllUserRequest struct {
	Page     int
	PageSize int
}

type UpdateUserRequest struct {
	ID       uint   `json:"id" validate:"required"`
	Name     string `json:"name,omitempty" validate:"max=100"`
	Password string `json:"password,omitempty" validate:"max=255"`
	Photo    string `json:"photo,omitempty" validate:"max=255"`
}
