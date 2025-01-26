package usecase

import (
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/repository"
	"github.com/go-playground/validator/v10"
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
