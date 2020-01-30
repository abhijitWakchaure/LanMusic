package logger

import (
	"log"
	"net/http"
	"time"
)

// LogLevel defines the severity of the logs
type LogLevel string

// const for different log levels
const (
	INFO     LogLevel = "[INFO]"
	DEBUG    LogLevel = "[DEBUG]"
	FATAL    LogLevel = "[FATAL]"
	CRITICAL LogLevel = "[CRITICAL]"
	DEFAULT  LogLevel = ""
)

// Log will do the actual logging
func Log(logLevel LogLevel, msg interface{}) {
	log.SetPrefix("[LanMusic] ")
	switch logLevel {
	case INFO:
		log.Println(INFO, msg)
	case DEBUG:
		log.Println(DEBUG, msg)
	case FATAL:
		log.Fatalln(FATAL, msg)
	case CRITICAL:
		log.Fatalln(CRITICAL, msg)
	default:
		log.Println(msg)
	}
}

// Logger ...
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
