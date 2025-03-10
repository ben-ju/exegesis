package main

import (
	"log"
	"net/http"

	"github.com/ben-ju/exegesis/internal/users"
)

func main() {
	// TODO : Change the os.Getenv to use the app.Config values add functions parameters and app methods to handle the call like GET, POST...
	// TODO: Initializing with file but might use grafana & prometheus
	app := mount()
	defer app.logger.File.Close()

	userRouter := http.NewServeMux()
	users.RegisterRoutes(userRouter)

	app.rootMux.Handle("/", userRouter)

	log.Fatal(app.run())
}
