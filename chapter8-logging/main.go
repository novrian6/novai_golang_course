package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	// Inisialisasi logger
	logger := logrus.New()

	// Menuliskan pesan log dengan level "Info"
	logger.Info("This is an info message")

	// Menuliskan pesan log dengan level "Warning"
	logger.Warn("This is a warning message")

	// Menuliskan pesan log dengan level "Error"
	logger.Error("This is an error message")
}

