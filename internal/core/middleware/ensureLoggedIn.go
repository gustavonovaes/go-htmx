package middleware

import (
	"log"
	"net/http"
)

func EnsureLoggedIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("INFO: EnsureLoggedIn middleware before request to %s", r.URL.String())

		// Get user from cookie
		// Check if user is logged in
		// If not, redirect to login page
		next.ServeHTTP(w, r)
	})
}
