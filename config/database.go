package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config *viper.Viper, log *logrus.Logger) *gorm.DB {
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(&loggerWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5, // Slow SQL threshold
			LogLevel:                  logger.Info,     // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,            // Don't include params in the SQL log
			Colorful:                  false,           // Disable color
		}),
	})
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
		panic(err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("connection failed: %v", err)
		panic(err)
	}

	connection.SetMaxIdleConns(maxIdleConnection)
	connection.SetMaxOpenConns(maxOpenConnection)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db
}

type loggerWriter struct {
	Logger *logrus.Logger
}

func (l *loggerWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
