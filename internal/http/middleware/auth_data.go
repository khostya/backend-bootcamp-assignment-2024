package middleware

import (
	"context"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/auth"
	"net/http"
	"strings"
)

const (
	KeyUserID   = "userID"
	KeyUserType = "user type"
)

func AuthData(manager auth.TokenManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenHeader := strings.Split(r.Header.Get("Authorization"), " ")

			userID, err := manager.ExtractUserId(tokenHeader)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), KeyUserID, userID))

			userType, err := manager.ExtractUserType(tokenHeader)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), KeyUserType, domain.UserType(userType)))

			next.ServeHTTP(w, r)
			return
		})
	}
}
