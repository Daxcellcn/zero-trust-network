package auth

import (
	"errors"
	"net/http"
)

// Dummy user store for example
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Authenticate user credentials
func Authenticate(username, password string) (bool, error) {
	if pwd, ok := users[username]; ok {
		if pwd == password {
			return true, nil
		}
		return false, errors.New("incorrect password")
	}
	return false, errors.New("user not found")
}

// Middleware for authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}
		authenticated, err := Authenticate(username, password)
		if err != nil || !authenticated {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
