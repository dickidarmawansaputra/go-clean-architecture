package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(config *viper.Viper, ctx context.Context) *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.Level(config.GetInt("LOG_LEVEL")))

	// set output log to log file
	file, err := os.OpenFile(fmt.Sprintf("./storage/logs/logger-%s.log", time.Now().Format(time.DateOnly)), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.WithContext(ctx).WithError(err).Fatal("failed to create log file")
	}

	log.SetOutput(file)

	return log
}
