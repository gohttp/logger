package logger

import "github.com/segmentio/go-log"
import "net/http"
import "time"

// wrapper to capture status.
type wrapper struct {
	http.ResponseWriter
	status int
}

// capture status.
func (w *wrapper) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// New logger middleware.
func New() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			res := &wrapper{w, 200}
			log.Info(">> %s %s", req.Method, req.RequestURI)
			h.ServeHTTP(res, req)
			log.Info("<< %s %s %d in %s", req.Method, req.RequestURI, res.status, time.Since(start))
		})
	}
}
