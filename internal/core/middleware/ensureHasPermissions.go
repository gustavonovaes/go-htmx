package middleware

import (
	"log"
	"net/http"
)

func EnsureHasPermissions(next http.Handler, permissions []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"INFO: EnsureHasPermissions middleware before request to %s with permissions %v",
			r.URL.String(),
			permissions,
		)

		// Get user from cookie
		// Check if user is logged in
		// Get user permissions
		// Check if user has permissions
		// If not, return 403 Forbidden
		next.ServeHTTP(w, r)
	})
}
