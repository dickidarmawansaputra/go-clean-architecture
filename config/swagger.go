package config

import (
	_ "github.com/dickidarmawansaputra/go-clean-architecture/docs"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
)

func NewSwagger(config *viper.Viper) *swagger.Config {
	return &swagger.Config{
		DeepLinking:  false,
		DocExpansion: "list",
	}
}
