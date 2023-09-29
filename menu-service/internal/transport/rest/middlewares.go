package rest

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	tokenUsers := map[string]string{
		"00000000": "admin",
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Auth-Token")

		if user, found := tokenUsers[token]; found {
			log.Printf("Authenticated user %s\n", user)
			next.ServeHTTP(w, r)
		}

		http.Error(w, "Forbidden", http.StatusForbidden)
	})
}
