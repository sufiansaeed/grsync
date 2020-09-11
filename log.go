package grsync

import (
	"os"

	logger "github.com/sirupsen/logrus"
)

func log() *logger.Entry {
	level, err := logger.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logger.SetLevel(logger.InfoLevel)
	}
	logger.SetLevel(level)

	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp: true,
	})

	return logger.WithFields(logger.Fields{
		"package": "grsync",
	})
}
