package main

import (
	"os"
	"github.com/sirupsen/logrus"
)

func main() {
	// Membuat file untuk menyimpan log
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Unable to create log file. Error:", err)
	}
	defer file.Close()

	// Inisialisasi logger dengan output ke file
	logrus.SetOutput(file)

	// Menuliskan pesan log dengan level "Info"
	logrus.Info("This is an info message")

	// Menuliskan pesan log dengan level "Warning"
	logrus.Warn("This is a warning message")

	// Menuliskan pesan log dengan level "Error"
	logrus.Error("This is an error message")
}

