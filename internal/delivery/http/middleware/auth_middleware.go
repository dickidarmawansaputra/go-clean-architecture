package middleware

import (
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/entity"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/exception"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/model"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/repository"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func NewAuthMiddleware(config *viper.Viper) func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.GetString("JWT_SECRET"))},
		ErrorHandler: jwtErrorHandler,
	})
}

func jwtErrorHandler(ctx *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return exception.Error(fiber.ErrBadRequest, err.Error())
	}

	return exception.Error(fiber.ErrUnauthorized, err.Error())
}

func AuthUser(ctx *fiber.Ctx, db *gorm.DB, repository *repository.UserRepository) (*model.UserResponse, error) {
	auth := ctx.Locals("user").(*jwt.Token)
	claims := auth.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	tx := db.Begin()
	defer tx.Rollback()

	user := new(entity.User)
	if err := repository.FindById(db, ctx, user, uint(id)); err != nil {
		return nil, err
	}

	return model.UserResource(ctx, user), nil
}
