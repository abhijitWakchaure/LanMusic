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
		log.Println(DEBUG, fmt.Sprintln(msg...))
	case INFO:
		log.Println(INFO, fmt.Sprintln(msg...))
	case ERROR:
		log.Println(ERROR, fmt.Sprintln(msg...))
	case FATAL:
		log.Fatalln(FATAL, fmt.Sprintln(msg...))
	case CRITICAL:
		log.Fatalln(CRITICAL, fmt.Sprintln(msg...))
	default:
		log.Println(fmt.Sprintln(msg...))
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
