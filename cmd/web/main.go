package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/dickidarmawansaputra/go-clean-architecture/config"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()
	log := config.NewLogger(cfg, ctx)
	app := config.NewFiber(cfg, log)

	config.Bootstrap(&config.BootstrapConfig{
		App:      app,
		Config:   cfg,
		DB:       config.NewDatabase(cfg, log),
		Validate: config.NewValidator(),
	})

	gracefulShutdown(ctx, app, cfg, log)
}

func gracefulShutdown(ctx context.Context, app *fiber.App, config *viper.Viper, log *logrus.Logger) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGHUP)
	defer stop()

	port := config.GetInt("APP_PORT")

	// listen from a different goroutine
	go func() {
		log.WithContext(ctx).Printf("starting server with graceful shutdown on: %d", port)
		err := app.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			log.WithContext(ctx).Fatalf("failed to start server: %v", err)
			panic(err)
		}
	}()

	<-ctx.Done()

	if err := app.Shutdown(); err != nil {
		log.WithContext(ctx).Fatalf("failed to graceful shutdown: %v", err)
		panic(err)
	}

	log.WithContext(ctx).Printf("fiber was successfully shutdown.")
}
