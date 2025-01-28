package config

import (
	_ "github.com/dickidarmawansaputra/go-clean-architecture/docs"
	"github.com/gofiber/swagger"
)

func NewSwagger() *swagger.Config {
	return &swagger.Config{
		DeepLinking:  false,
		DocExpansion: "list",
	}
}
