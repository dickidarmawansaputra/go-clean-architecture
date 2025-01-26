package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config *viper.Viper) *gorm.DB {
	host := config.GetString("DB_HOST")
	username := config.GetString("DB_USER")
	password := config.GetString("DB_PASSWORD")
	database := config.GetString("DB_NAME")
	port := config.GetInt("DB_PORT")
	maxIdleConnection := config.GetInt("DB_MAX_IDLE")
	maxOpenConnection := config.GetInt("DB_MAX_OPEN")
	maxLifeTimeConnection := config.GetInt("DB_MAX_LIFETIME")
	timezone := config.GetString("DB_TIMEZONE")
	sslmode := config.GetString("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", host, username, password, database, port, sslmode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	connection, err := db.DB()
	if err != nil {
		panic(err)
	}

	connection.SetMaxIdleConns(maxIdleConnection)
	connection.SetMaxOpenConns(maxOpenConnection)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db
}
