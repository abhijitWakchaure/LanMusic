package logger

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// LogLevel defines the severity of the logs
type LogLevel string

// const for different log levels
const (
	DEBUG    LogLevel = "[DEBUG]"
	INFO     LogLevel = "[INFO]"
	ERROR    LogLevel = "[ERROR]"
	FATAL    LogLevel = "[FATAL]"
	CRITICAL LogLevel = "[CRITICAL]"
	DEFAULT  LogLevel = ""
)

func init() {
	log.SetPrefix("# ")
}

// Log will do the actual logging
func Log(logLevel LogLevel, msg ...interface{}) {
	switch logLevel {
	case DEBUG:
		log.Print(DEBUG, fmt.Sprintln(msg...))
	case INFO:
		log.Print(INFO, fmt.Sprintln(msg...))
	case ERROR:
		log.Print(ERROR, fmt.Sprintln(msg...))
	case FATAL:
		log.Fatal(FATAL, fmt.Sprintln(msg...))
	case CRITICAL:
		log.Fatal(CRITICAL, fmt.Sprintln(msg...))
	default:
		log.Print(fmt.Sprintln(msg...))
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
