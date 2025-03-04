package middleware

import (
	"net/http"
)

var DEFAULT_MIDDLEWARES = []func(http.HandlerFunc) http.HandlerFunc{
	recovery,
	logging,
}

func SetMiddlewares(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	h := handler
	for _, mw := range DEFAULT_MIDDLEWARES {
		h = mw(h)
	}
	for _, mw := range middlewares {
		h = mw(h)
	}
	return h
}
