package main

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs each incoming request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		// BEFORE handler
		log.Printf("Incoming request → %s %s", r.Method, r.URL.Path)

		// Call next handler
		next.ServeHTTP(w, r)

		// AFTER handler
		log.Printf("Completed → %s %s (%v)", r.Method, r.URL.Path, time.Since(start))
	})
}
