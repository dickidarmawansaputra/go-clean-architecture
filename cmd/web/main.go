package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/dickidarmawansaputra/go-clean-architecture/config"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	cfg := config.NewConfig()
	app := config.NewFiber(cfg)

	config.Bootstrap(&config.BootstrapConfig{
		App:      app,
		Config:   cfg,
		DB:       config.NewDatabase(cfg),
		Validate: config.NewValidator(),
	})

	gracefulShutdown(context.Background(), app, cfg)
}

func gracefulShutdown(ctx context.Context, app *fiber.App, config *viper.Viper) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGHUP)
	defer stop()

	port := config.GetInt("APP_PORT")

	// listen from a different goroutine
	go func() {
		err := app.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			panic(err)
		}
	}()

	<-ctx.Done()

	if err := app.Shutdown(); err != nil {
		panic(err)
	}
}
