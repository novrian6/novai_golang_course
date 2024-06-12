package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	// Create a new logger instance
	logger := logrus.New()

	// Log in JSON format for structured logging
	logger.SetFormatter(&logrus.JSONFormatter{})

	// You can set log level to control verbosity
	logger.SetLevel(logrus.InfoLevel)

	// You can also set the output destination, like a file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}

	// Example of structured logging
	logger.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logger.WithFields(logrus.Fields{
		"temperature": -4,
	}).Error("It's too cold")

	logger.WithFields(logrus.Fields{
		"user": "john",
		"age":  30,
	}).Warn("User is getting old")

	logger.Info("Logging complete")
}
