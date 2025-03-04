package main

import (
	"database/sql"
	"net/http"

	"github.com/ben-ju/exegesis/internal/config"
	"github.com/ben-ju/exegesis/internal/middleware"
	"github.com/ben-ju/exegesis/internal/service"
	"github.com/ben-ju/exegesis/internal/utils"
)

type App struct {
	DB          *sql.DB
	Server      *http.Server
	Middlewares []func(http.HandlerFunc) http.HandlerFunc
	Logger      *utils.Logger
	Router      *http.ServeMux
	Config      *config.Config
}

func NewApp() *App {
	db := service.InitDB()
	config := config.NewConfig()
	rootMux := http.NewServeMux()
	server := newHTTPServer(rootMux)
	logger := utils.NewLogger()
	app := &App{
		DB:     db,
		Server: server,
		Logger: logger,
		Router: rootMux,
		Config: config,
	}
	app.initDefaultMiddlewares()
	return app
}

func (a *App) initDefaultMiddlewares() {
	a.Middlewares = append(a.Middlewares, middleware.DEFAULT_MIDDLEWARES...)
}
