package utility

import (
	"time"

	"github.com/sirupsen/logrus"
)

func InitializeLogger() *logrus.Entry {
	var Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})

	Logger.SetLevel(logrus.TraceLevel)

	log := Logger.WithFields(logrus.Fields{
		"port":          8080,
		"database_port": "PostgreSQL: 54321",
		"database_host": "PostgreSQL: localhost",
		"status":        "running",
	})
	log.Info("Logger initialized")

	return log
}
