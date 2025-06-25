package helpers

import "github.com/sirupsen/logrus"

func SetupLogger() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	log.Info("logger initiated using logrus")

	return log
}
