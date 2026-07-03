// A response time middleware will calculate the time it took to process a request from a client and send a response

// It is going to start tracking time as soon as we receive the request and it is going to calculate the time it took to send a response back to the client
package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Recevied Request in ResponseTime")
		start := time.Now()

		// We have to create a custom response writer to capture the status code, because we are also going to capture the status code and we are going to log that

		wrappedWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		
		
		duration := time.Since(start)
		w.Header().Set("X-Response-Time", duration.String())
		next.ServeHTTP(wrappedWriter, r)

		// Calculate the Duration
		duration = time.Since(start)
		// log the request details
		fmt.Printf("Method: %s, URL: %s, Status: %d, Duration: %v\n", r.Method, r.URL, wrappedWriter.status, duration.String())
		fmt.Println("Sent Response from Response Time Middleware")

	})
}

// response Writer
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
