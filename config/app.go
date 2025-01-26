package config

import (
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/delivery/controller"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/delivery/http/route"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/repository"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App      *fiber.App
	Config   *viper.Viper
	DB       *gorm.DB
	Validate *validator.Validate
}

func Bootstrap(config *BootstrapConfig) {
	// repositories
	userRepository := repository.NewUserRepository()

	// usercases
	authUseCase := usecase.NewAuthUseCase(config.Config, config.DB, config.Validate, userRepository)

	// route config
	routeConfig := route.RouteConfig{
		App:            config.App,
		AuthController: controller.NewAuthController(authUseCase),
	}

	route.Router(&routeConfig)
}
