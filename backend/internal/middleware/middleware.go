package middleware

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"
)

var DEFAULT_MIDDLEWARES = []func(http.HandlerFunc) http.HandlerFunc{
	Recovery,
	Logging,
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

// Logging is a middleware that logs the details of every request using a default configuration.
func Logging(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler in the chain.
		next(w, r)

		// Build log message with all desired fields.
		logMsg := fmt.Sprintf(
			"\n\n[REQUEST LOG]\n"+
				"Time: %v\n"+
				"Method: %v\n"+
				"Host: %v\n"+
				"Route: %v\n"+
				"Path: %v\n"+
				"Referer: %v\n"+
				"IP Address: %v\n"+
				"User Agent: %v\n"+
				"Request Duration: %.3fs\n",
			time.Now().Format("2006-01-02 15:04:05"), // Current time.
			r.Method,                                 // HTTP method.
			r.Host,                                   // Request host.
			r.URL.Path,                               // Route (URL path).
			r.URL.RequestURI(),                       // Full request URI.
			r.Referer(),                              // Referer header.
			r.RemoteAddr,                             // Remote IP address.
			r.UserAgent(),                            // User agent.
			time.Since(start).Seconds(),              // Request duration.
		)

		log.Printf("%v", logMsg)
	})
}

func Recovery(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next(w, r)
	}
}

func Test(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("1")
		next(w, r)
		log.Println("4")
	})
}

func SecondTest(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("2")
		next(w, r)
		log.Println("3")
	})
}
