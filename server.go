package server

import (
	"net/http"
	"zero-trust-network/auth"   // This should be relative to the module path
	"zero-trust-network/policy" // This should be relative to the module path
)

func StartServer() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/resource1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Resource 1"))
	})
	mux.HandleFunc("/resource2", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Resource 2"))
	})

	handler := auth.AuthMiddleware(policy.PolicyMiddleware(mux))

	return http.ListenAndServe(":8080", handler)
}
