package main

import (
	"log"
	"net/http"

	"github.com/ben-ju/exegesis/internal/repository"
	"github.com/ben-ju/exegesis/internal/router"
	"github.com/ben-ju/exegesis/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	// TODO : Change the os.Getenv to use the app.Config values add functions parameters and app methods to handle the call like GET, POST...
	// TODO: Initializing with file but might use grafana & prometheus
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env : %v", err)
	}
	app := mount()
	defer app.logger.File.Close()
	defer app.db.Close()

	bookRepository := repository.NewBookRepository(app.db)
	bookService := service.NewBookRoutable(bookRepository)
	bookRouter := router.NewRouter()
	// booksRouter.Use(middleware.Test, middleware.SecondTest)
	bookRouter.Register(bookService)

	app.rootMux.Handle("/books/", http.StripPrefix("/books", bookRouter))

	log.Fatal(app.run())
}
