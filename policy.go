package policy

import (
	"errors"
	"net/http"
)

// Define access policies
var policies = map[string][]string{
	"user1": {"/resource1", "/resource2"},
	"user2": {"/resource2"},
}

// Check if user has access to the resource
func CheckAccess(username, resource string) (bool, error) {
	allowedResources, ok := policies[username]
	if !ok {
		return false, errors.New("policy not found for user")
	}
	for _, r := range allowedResources {
		if r == resource {
			return true, nil
		}
	}
	return false, errors.New("access denied")
}

// Middleware for authorization
func PolicyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, _, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}
		resource := r.URL.Path
		allowed, err := CheckAccess(username, resource)
		if err != nil || !allowed {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
