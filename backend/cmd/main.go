package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ben-ju/exegesis/internal/middleware"
	"github.com/ben-ju/exegesis/internal/utils"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world !")
}

func main() {
	// Initializing with file but might use grafana & prometheus
	logFile := utils.InitLogFile()
	defer logFile.Close()

	h := middleware.SetMiddlewares(handler, middleware.Logging())
	http.HandleFunc("/", h)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), nil))
}
