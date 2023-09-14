package logger

import "github.com/sirupsen/logrus"

var logger *logrus.Logger

func NewLogger() *logrus.Logger {
	if logger == nil {
		logger = logrus.New()
		logger.SetFormatter(&logrus.JSONFormatter{})
	}

	return logger
}
