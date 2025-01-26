package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      config.GetString("APP_NAME"),
		ErrorHandler: errorHandler(),
		Prefork:      config.GetBool("APP_PREFORK"),
	})

	return app
}

func errorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		var errors any
		status := fiber.ErrInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			status = e
			errors = e.Error()
		}

		return ctx.Status(status.Code).JSON(fiber.Map{
			"code":   status.Code,
			"status": status.Message,
			"errors": errors,
		})
	}
}
