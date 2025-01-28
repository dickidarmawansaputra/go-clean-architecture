package repository

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository[T any] struct{}

func (r *Repository[T]) Create(db *gorm.DB, ctx *fiber.Ctx, entity *T) error {
	return db.WithContext(ctx.UserContext()).Create(entity).Error
}

func (r *Repository[T]) FindById(db *gorm.DB, ctx *fiber.Ctx, entity *T, id uint) error {
	return db.WithContext(ctx.UserContext()).Where("id = ?", id).Take(entity).Error
}

func (r *Repository[T]) Update(db *gorm.DB, ctx *fiber.Ctx, entity *T, id uint) error {
	return db.WithContext(ctx.UserContext()).Where("id = ?", id).Save(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, ctx *fiber.Ctx, entity *T, id uint) error {
	return db.WithContext(ctx.UserContext()).Where("id = ?", id).Delete(entity).Error
}

func (r *Repository[T]) CountByField(db *gorm.DB, ctx *fiber.Ctx, entity *T, field string, value any) (int64, error) {
	var total int64
	err := db.WithContext(ctx.UserContext()).Model(&entity).Where(fmt.Sprintf("%s = ?", field), value).Count(&total).Error

	return total, err
}

func (r *Repository[T]) FindByField(db *gorm.DB, ctx *fiber.Ctx, entity *T, field string, value any) error {
	return db.WithContext(ctx.UserContext()).Where(fmt.Sprintf("%s = ?", field), value).Take(&entity).Error
}
