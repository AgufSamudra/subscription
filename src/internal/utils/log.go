package utils

import (
	stdlog "log"
	"os"
)

var logger = stdlog.New(os.Stdout, "", stdlog.Ldate|stdlog.Ltime)

func Info(message string) {
	logger.Printf("[INFO] %s", message)
}

func Infof(format string, args ...any) {
	logger.Printf("[INFO] "+format, args...)
}

func Error(message string) {
	logger.Printf("[ERROR] %s", message)
}

func Errorf(format string, args ...any) {
	logger.Printf("[ERROR] "+format, args...)
}

func Fatal(message string) {
	logger.Fatalf("[ERROR] %s", message)
}

func Fatalf(format string, args ...any) {
	logger.Fatalf("[ERROR] "+format, args...)
}
