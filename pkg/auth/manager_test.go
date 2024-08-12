package auth

import (
	"backend-bootcamp-assignment-2024/internal/domain"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestManager(t *testing.T) {
	t.Parallel()

	user := domain.User{
		ID:       uuid.New(),
		Email:    "radium@radium-rtf.ru",
		UserType: domain.UserClient,
	}

	tests := []struct {
		manager TokenManager
		name    string
		user    domain.User
		ttl     time.Duration
		err     error
	}{
		{
			manager: NewManager("arawdawdwea"),
			name:    "ok",
			user:    user,
			ttl:     time.Hour,
		},
		{
			manager: NewManager("arawdawdwea"),
			name:    "exp",
			user:    user,
			ttl:     0,
			err:     jwt.ErrTokenExpired,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			user := tt.user
			token, err := tt.manager.NewUserJWT(user.ID, string(user.UserType), time.Now().Add(tt.ttl))
			require.NoError(t, err)

			userType, err := tt.manager.ExtractUserType([]string{"Bearer", token})
			if errors.Is(err, tt.err) {
				return
			}

			require.NoError(t, err)
			require.Equal(t, string(user.UserType), userType)

			userID, err := tt.manager.ExtractUserId([]string{"Bearer", token})
			require.NoError(t, err)
			require.Equal(t, userID, user.ID)
		})
	}
}
