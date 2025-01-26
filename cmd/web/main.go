package main

import "github.com/dickidarmawansaputra/go-clean-architecture/config"

func main() {
	cfg := config.NewConfig()
	app := config.NewFiber(cfg)

	config.Bootstrap(&config.BootstrapConfig{
		App:      app,
		Config:   cfg,
		DB:       config.NewDatabase(cfg),
		Validate: config.NewValidator(),
	})
}
