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

// @title			Go Clean Architecture With Fiber
// @version		1.0.0
// @description	Go Clean Architecture With Fiber Framework
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	fiber@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
// @BasePath		/
// @securityDefinitions.apikey  Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	ctx := context.Background()
	cfg := config.NewConfig()
	log := config.NewLogger(cfg, ctx)
	app := config.NewFiber(cfg, log)

	config.Bootstrap(&config.BootstrapConfig{
		App:      app,
		Config:   cfg,
		DB:       config.NewDatabase(ctx, cfg, log),
		Validate: config.NewValidator(),
		Swagger:  config.NewSwagger(),
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

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.WithContext(ctx).Fatalf("failed to graceful shutdown: %v", err)
		panic(err)
	}

	log.WithContext(ctx).Printf("fiber was successfully shutdown.")
}
