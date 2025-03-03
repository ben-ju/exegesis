package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type LoggerConfig struct {
	LoggerFormat LoggerFormatOptions
}

type LoggerFormatOptions struct {
	UserAgent *string    `json:"user_agent,omitempty"`
	Time      *time.Time `json:"time,omitempty"`
	RequestId *string    `json:"request_id,omitempty"`
	Method    *string    `json:"method,omitempty"`
	Host      *string    `json:"host,omitempty"`
	Route     *string    `json:"route,omitempty"`
	Path      *string    `json:"path,omitempty"`
	Referer   *string    `json:"referer,omitempty"`
	IpAddress *string    `json:"ip_address,omitempty"`
}

var DefaultLoggerConfig = LoggerConfig{
	LoggerFormat: LoggerFormatOptions{
		UserAgent: new(string),
		Time:      new(time.Time),
		RequestId: new(string),
		Method:    new(string),
		Host:      new(string),
		Route:     new(string),
		Path:      new(string),
		Referer:   new(string),
		IpAddress: new(string),
	},
}

func Logging() MiddlewareFunc {
	return LoggingWithConfig(DefaultLoggerConfig)
}

func LoggingWithConfig(loggerConfig LoggerConfig) MiddlewareFunc {
	return func(next HandlerFunc) HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			next(w, r)

			logData := make(map[string]interface{})
			for field, value := range map[string]*string{
				"user_agent": loggerConfig.LoggerFormat.UserAgent,
				"request_id": loggerConfig.LoggerFormat.RequestId,
				"method":     loggerConfig.LoggerFormat.Method,
				"host":       loggerConfig.LoggerFormat.Host,
				"route":      loggerConfig.LoggerFormat.Route,
				"path":       loggerConfig.LoggerFormat.Path,
				"referer":    loggerConfig.LoggerFormat.Referer,
				"ip_address": loggerConfig.LoggerFormat.IpAddress,
			} {
				if value != nil {
					switch field {
					case "user_agent":
						logData[field] = r.UserAgent()
					case "request_id":
						logData[field] = r.Header.Get("X-Request-ID")
					case "method":
						logData[field] = r.Method
					case "host":
						logData[field] = r.Host
					case "route":
						logData[field] = r.URL.Path
					case "path":
						logData[field] = r.URL.RequestURI()
					case "referer":
						logData[field] = r.Referer()
					case "ip_address":
						logData[field] = r.RemoteAddr
					}
				}
			}
			if loggerConfig.LoggerFormat.Time != nil {
				logData["time"] = time.Now().Format("2006-01-02 15:04:05")
			}
			duration := time.Since(start).Seconds()
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
				logData["time"],
				logData["method"],
				logData["host"],
				logData["route"],
				logData["path"],
				logData["referer"],
				logData["ip_address"],
				logData["user_agent"],
				duration,
			)

			log.Printf("%v", logMsg)
		}
	}
}
