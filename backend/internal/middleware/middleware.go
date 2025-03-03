package middleware

import "net/http"

// HandlerFunc defines the type for HTTP handlers
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// MiddlewareFunc defines the type for middleware functions
type MiddlewareFunc func(next HandlerFunc) HandlerFunc
