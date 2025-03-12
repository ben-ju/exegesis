package main

import (
	"log"

	"github.com/ben-ju/exegesis/internal/middleware"
	"github.com/ben-ju/exegesis/internal/router"
	"github.com/ben-ju/exegesis/internal/users"
)

func main() {
	// TODO : Change the os.Getenv to use the app.Config values add functions parameters and app methods to handle the call like GET, POST...
	// TODO: Initializing with file but might use grafana & prometheus
	app := mount()
	defer app.logger.File.Close()

	userRouter := router.NewRouter()

	userRouter.Use(middleware.Test, middleware.SecondTest)
	userRouter.Register(&users.UserRoutable{})

	app.rootMux.Handle("/", userRouter)

	log.Fatal(app.run())
}
