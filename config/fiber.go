package config

import (
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/exception"
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper, log *logrus.Logger) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      config.GetString("APP_NAME"),
		ErrorHandler: errorHandler(log),
		Prefork:      config.GetBool("APP_PREFORK"),
	})

	return app
}

func errorHandler(log *logrus.Logger) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		var errors any
		status := fiber.ErrInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			status = e
			errors = e.Error()
		}

		if e, ok := err.(*exception.ErrorResponse); ok {
			status = e.Status
			errors = e.Errors
		}

		log.WithContext(ctx.UserContext()).Log(log.GetLevel(), errors)

		return model.ErrorResponse(ctx, status, errors)
	}
}
