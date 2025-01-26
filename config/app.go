package config

import (
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/delivery/http/route"
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
	// route config
	routeConfig := route.RouteConfig{
		App: config.App,
	}

	route.Router(&routeConfig)
}
