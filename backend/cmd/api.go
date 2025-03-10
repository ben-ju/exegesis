package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/ben-ju/exegesis/internal/config"
	"github.com/ben-ju/exegesis/internal/service"
)

type app struct {
	db      *sql.DB
	server  *http.Server
	logger  *service.Logger
	rootMux *http.ServeMux
	config  config.Config
}

func mount() *app {
	db := service.InitDB()
	config := config.NewConfig()
	port := config.AppPort
	if port == "" {
		port = "8080"
	}
	rootMux := http.NewServeMux()

	server := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    time.Minute,
		Handler:        rootMux,
		MaxHeaderBytes: 1 << 20,
	}
	logger := service.NewLogger()
	app := &app{
		db:      db,
		server:  server,
		logger:  logger,
		rootMux: rootMux,
		config:  config,
	}
	return app
}

func (a *app) run() error {
	return a.server.ListenAndServe()
}
