package utils

import (
	"log"
	"os"
)

func InitLog() *os.File {
	f, err := os.OpenFile("go-backend.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Panic(err)
	}
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(f)
	return f
}
