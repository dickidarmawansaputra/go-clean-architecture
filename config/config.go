package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewConfig() *viper.Viper {
	config := viper.New()
	config.AddConfigPath(".")
	config.SetConfigFile(".env")

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	return config
}
