package middleware

import (
	"log"
	"net/http"
	"runtime"
	"time"
)

var m runtime.MemStats

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		runtime.ReadMemStats(&m)
		log.Printf(
			"INFO: %s %v %s - %.2f MB",
			r.Method,
			time.Since(start),
			r.URL.Path,
			float64(m.Alloc)/1024/1024,
		)
	})
}
