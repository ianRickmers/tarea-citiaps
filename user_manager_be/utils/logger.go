package utils

import (
	"log"
	"time"
)

// Logger es para manejar el registro en la aplicación

// Info registra un mensaje de información
func Info(msg string) {
	log.Printf("[INFO] %s - %s\n", time.Now().Format(time.RFC3339), msg)
}

// Error registra un mensaje de error
func Error(err error, context string) {
	log.Printf("[ERROR] %s - %s: %v\n", time.Now().Format(time.RFC3339), context, err)
}
