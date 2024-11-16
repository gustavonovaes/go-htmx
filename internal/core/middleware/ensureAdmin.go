package middleware

import (
	"log"
	"net/http"
)

func EnsureAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("INFO: EnsureAdmin middleware before request to %s", r.URL.String())

		// Get user from cookie
		// Check if user is admin
		// If not, return 403 Forbidden
		next.ServeHTTP(w, r)
	})
}
