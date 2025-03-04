package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Logging is a middleware that logs the details of every request using a default configuration.
func logging(next http.HandlerFunc) http.HandlerFunc {
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
