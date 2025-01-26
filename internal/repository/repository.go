package repository

import "gorm.io/gorm"

type Repository[T any] struct{}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) FindById(db *gorm.DB, entity *T, id uint) error {
	return db.Where("id = ?", id).Take(entity).Error
}

func (r *Repository[T]) Update(db *gorm.DB, entity *T, id uint) error {
	return db.Where("id = ?", id).Save(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entity *T, id uint) error {
	return db.Where("id = ?", id).Delete(entity).Error
}
