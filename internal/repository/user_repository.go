package repository

import (
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/entity"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CheckUserExists(db *gorm.DB, ctx *fiber.Ctx, entity *entity.User, email string) bool {
	total, _ := r.CountByField(db, ctx, entity, "email", email)
	if total > 0 {
		return true
	}

	return false
}
