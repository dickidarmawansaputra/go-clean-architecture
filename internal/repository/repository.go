package repository

import (
	"fmt"
	"math"

	"github.com/dickidarmawansaputra/go-clean-architecture/internal/model"
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

func (r *Repository[T]) Count(db *gorm.DB, ctx *fiber.Ctx, entity *T) (int64, error) {
	var total int64
	err := db.WithContext(ctx.UserContext()).Model(entity).Count(&total).Error

	return total, err
}

func (r *Repository[T]) Paginate(db *gorm.DB, ctx *fiber.Ctx, entity *[]T, page int, pageSize int) ([]T, *model.MetaPagination, error) {
	offset := (page - 1) * pageSize
	err := db.WithContext(ctx.UserContext()).Limit(pageSize).Offset(offset).Find(&entity).Error
	if err != nil {
		return nil, nil, err
	}

	total, _ := r.Count(db, ctx, new(T))
	count := len(*entity)
	totalPage := int(math.Ceil(float64(total / int64(pageSize))))

	var nextPage int
	var prevPage int
	if count < pageSize {
		nextPage = 0
		prevPage = totalPage
	} else {
		nextPage = page + 1
		prevPage = page - 1
	}

	var nextLink string
	if nextPage != 0 {
		nextLink = fmt.Sprintf("%s%s?page=%d&page_size=%d", ctx.BaseURL(), ctx.Route().Path, nextPage, pageSize)
	}

	var prevLink string
	if prevPage != 0 {
		prevLink = fmt.Sprintf("%s%s?page=%d&page_size=%d", ctx.BaseURL(), ctx.Route().Path, prevPage, pageSize)
	}

	meta := &model.PaginationMetaData{
		Total:       int(total),
		Count:       count,
		PerPage:     pageSize,
		CurrentPage: page,
		TotalPage:   totalPage,
		Links: &model.PaginationLink{
			NextPage:     nextLink,
			PreviousPage: prevLink,
		},
	}

	return *entity, &model.MetaPagination{Pagination: meta}, nil
}
