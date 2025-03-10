package middleware

import (
	"net/http"
	"reflect"
)

var DEFAULT_MIDDLEWARES = []func(http.HandlerFunc) http.HandlerFunc{
	recovery,
	logging,
}
// This function is responsible for setting middlewares unto a handler (including the default). It is also preventing from using duplicate middlewares
func SetMiddlewares(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {

	uniqueMap := make(map[uintptr]struct{})
	var result []func(http.HandlerFunc) http.HandlerFunc

	addMiddleware := func(mw func(http.HandlerFunc) http.HandlerFunc) {
		ptr := reflect.ValueOf(mw).Pointer()
		if _, ok := uniqueMap[ptr]; !ok {
			uniqueMap[ptr] = struct{}{}
			result = append(result, mw)
		}
	}

	for _, mw := range DEFAULT_MIDDLEWARES {
		addMiddleware(mw)
	}
	for _, mw := range middlewares {
		addMiddleware(mw)
	}
	for _, mw := range result {
		handler = mw(handler)
	}

    return handler
}
