package main

import (
	"log"
	"os"
	"time"
)

// CustomLogger is a custom logger with different log levels.
type CustomLogger struct {
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
}

// NewCustomLogger creates a new CustomLogger.
func NewCustomLogger(infoHandle, warningHandle, errorHandle *os.File) *CustomLogger {
	return &CustomLogger{
		infoLog:    log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warningLog: log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog:   log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info logs an informational message.
func (l *CustomLogger) Info(message string) {
	l.infoLog.Println(message)
}

// Warning logs a warning message.
func (l *CustomLogger) Warning(message string) {
	l.warningLog.Println(message)
}

// Error logs an error message.
func (l *CustomLogger) Error(message string) {
	l.errorLog.Println(message)
}

func main() {
	// Create log files
	infoFile, err := os.OpenFile("info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer infoFile.Close()

	warningFile, err := os.OpenFile("warning.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer warningFile.Close()

	errorFile, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer errorFile.Close()

	// Create a new custom logger
	logger := NewCustomLogger(infoFile, warningFile, errorFile)

	// Log some messages
	logger.Info("This is an info message")
	logger.Warning("This is a warning message")
	logger.Error("This is an error message")

	// Simulate a long-running process
	for i := 0; i < 5; i++ {
		logger.Info("Processing step " + string(i))
		time.Sleep(1 * time.Second)
	}

	logger.Info("Process completed")
}
