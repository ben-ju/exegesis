package service

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	Instance *log.Logger
	File     *os.File
}

const LOG_FILE = "go-backend.log"

func NewLogger() *Logger {
	file, err := os.OpenFile(LOG_FILE, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Panicf("Failed to open log file: %v", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	instance := log.New(multiWriter, "APP_LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &Logger{
		Instance: instance,
		File:     file,
	}
}
