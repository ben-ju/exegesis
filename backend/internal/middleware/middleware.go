package middleware

import "net/http"

// MiddlewareFunc defines the type for middleware functions

func SetMiddlewares(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	h := handler
	h = recovery(h)
	for _, mw := range middlewares {
		h = mw(h)
	}
	return h
}
