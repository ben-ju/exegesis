package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ben-ju/exegesis/internal/middleware"
	"github.com/ben-ju/exegesis/internal/utils"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
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
	// TODO: Initializing with file but might use grafana & prometheus
	logFile := utils.InitLogFile()
	defer logFile.Close()
	rootMux := http.NewServeMux()

	homeHandler := middleware.SetMiddlewares(home, middleware.Logging())
	loginHandler := middleware.SetMiddlewares(login, middleware.Logging())
	registerHandler := middleware.SetMiddlewares(register, middleware.Logging())
	logoutHandler := middleware.SetMiddlewares(logout, middleware.Logging())

	webMux := http.NewServeMux()
	webMux.HandleFunc("/", homeHandler)
	webMux.HandleFunc("/login", loginHandler)
	webMux.HandleFunc("POST /register", registerHandler)

	userMux := http.NewServeMux()
	userMux.HandleFunc("/logout", logoutHandler)

	rootMux.Handle("/", webMux)
	rootMux.Handle("/user", userMux)

	server := newHTTPServer(rootMux)
	log.Fatal(server.ListenAndServe())
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
