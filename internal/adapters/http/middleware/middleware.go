package middleware

import (
	"context"
	"net/http"
	"strings"

	"meetup-app-hexa-arch/internal/adapters/auth"

	"github.com/dgrijalva/jwt-go"
)

type AuthMiddleware struct {
	jwtService *auth.JWTService
}

func NewAuthMiddleware(jwtService *auth.JWTService) *AuthMiddleware {
	return &AuthMiddleware{jwtService: jwtService}
}

func (m *AuthMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract Authorization header
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		if tokenString == "" {
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}

		// Validate the token
		token, err := m.jwtService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Extract claims and pass the userID in context
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unable to parse token claims", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["sub"].(string)
		if !ok || userID == "" {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Add userID to the request context
		ctx := context.WithValue(r.Context(), "userID", userID)
		r = r.WithContext(ctx)

		// Pass control to the next handler
		next.ServeHTTP(w, r)
	})
}
