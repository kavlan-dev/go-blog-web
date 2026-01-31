package middleware

import (
	"go-blog-web/internal/model"
	"net/http"
)

type userService interface {
	AuthenticateUser(username, password string) (*model.User, error)
}

func AuthMiddleware(service userService, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name, pass, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, err := service.AuthenticateUser(name, pass)
		if err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
