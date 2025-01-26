package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID              uint         `gorm:"primarykey"`
	Name            string       `gorm:"column:name"`
	Email           string       `gorm:"column:email"`
	Photo           string       `gorm:"column:photo"`
	EmailVerifiedAt sql.NullTime `gorm:"column:email_verified_at"`
	Password        string       `gorm:"column:password"`
	CreatedAt       time.Time    `gorm:"column:created_at"`
	UpdatedAt       time.Time    `gorm:"column:updated_at"`
}

func (u *User) TableName() string {
	return "users"
}
