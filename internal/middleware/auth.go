package middleware

import (
	"net/http"
	"strings"

	"your_project/internal/utils"

	"github.com/gorilla/context"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, _ := token.Claims.(map[string]interface{})
		context.Set(r, "user_email", claims["email"])

		next.ServeHTTP(w, r)
	})
}
