package repository

import "github.com/dickidarmawansaputra/go-clean-architecture/internal/entity"

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
