package logger

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	Body       bytes.Buffer
}

func LogRequest(r *http.Request) {
	log.Printf(
		"[REQ] %s %s | From: %s | UA: %s | Time: %s",
		r.Method,
		r.URL.Path,
		r.RemoteAddr,
		r.UserAgent(),
		time.Now().Format(time.RFC3339),
	)
}

func LogResponse(status int, size int) {
	log.Printf(
		"[RES] Status: %d | Size: %d bytes | Time: %s",
		status,
		size,
		time.Now().Format(time.RFC3339),
	)
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
	}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *LoggingResponseWriter) Write(b []byte) (int, error) {
	lrw.Body.Write(b) // captura o body da response
	return lrw.ResponseWriter.Write(b)
}
