package router

import (
	"net/http"
)

type Routable interface {
	RegisterRoutes(r *Router)
}

type Router struct {
	*http.ServeMux
	middlewares []func(http.HandlerFunc) http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		ServeMux:    http.NewServeMux(),
		middlewares: []func(http.HandlerFunc) http.HandlerFunc{},
	}
}

func (r *Router) Register(routable Routable) {
	routable.RegisterRoutes(r)
}

func (r *Router) Use(mws ...func(http.HandlerFunc) http.HandlerFunc) {
	r.middlewares = append(r.middlewares, mws...)
}

func (r *Router) HandleFunc(pattern string, handler http.HandlerFunc) {
	wrapped := handler

	// Applying middlewares in the reverse order
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		wrapped = r.middlewares[i](wrapped)
	}

	r.ServeMux.HandleFunc(pattern, wrapped)
}
