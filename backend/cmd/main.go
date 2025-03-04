package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ben-ju/exegesis/internal/middleware"
)

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Login Page")
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "register page")
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "logout page")
}

func main() {
	// TODO : Change the os.Getenv to use the app.Config values add functions parameters and app methods to handle the call like GET, POST...
	// TODO: Initializing with file but might use grafana & prometheus
	app := NewApp()
	defer app.Logger.File.Close()

	// TODO : Change the setMiddlewares so it's not mandatory even when there is no middleware (except default)
	homeHandler := middleware.SetMiddlewares(home)
	loginHandler := middleware.SetMiddlewares(login)
	registerHandler := middleware.SetMiddlewares(register)
	logoutHandler := middleware.SetMiddlewares(logout)

	webMux := http.NewServeMux()
	webMux.HandleFunc("/", homeHandler)
	webMux.HandleFunc("/login", loginHandler)
	webMux.HandleFunc("POST /register", registerHandler)

	userMux := http.NewServeMux()
	userMux.HandleFunc("/logout", logoutHandler)

	app.Router.Handle("/", webMux)
	app.Router.Handle("/user", userMux)

	log.Fatal(app.Server.ListenAndServe())
}

func newHTTPServer(rootMux *http.ServeMux) *http.Server {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	return &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        rootMux,
		MaxHeaderBytes: 1 << 20,
	}
}
