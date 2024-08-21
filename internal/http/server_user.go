package http

import (
	"context"
	"github.com/google/uuid"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/middleware"
	"net/http"
)

type (
	userUseCase interface {
		GetByID(ctx context.Context, id uuid.UUID) (domain.User, error)
	}
)

func (s *server) GetUser(w http.ResponseWriter, r *http.Request) {
	type resp struct {
		Email    string    `json:"email"`
		UserId   uuid.UUID `json:"user_id"`
		UserType string    `json:"user_type"`
	}

	status, user, err := s.getUser(r.Context())
	if err != nil {
		s.error(w, r, status, err)
		return
	}

	s.json(w, r, http.StatusOK, resp{Email: user.Email, UserType: string(user.UserType), UserId: user.ID})
}

func (s *server) getUser(ctx context.Context) (int, domain.User, error) {
	isDummy, ok := ctx.Value(middleware.KeyIsDummy).(bool)
	if !ok || isDummy {
		return http.StatusUnauthorized, domain.User{}, errUnauthorized
	}

	userID := ctx.Value(middleware.KeyUserID).(uuid.UUID)
	user, err := s.useCases.User.GetByID(ctx, userID)
	if err != nil {
		return http.StatusInternalServerError, domain.User{}, err
	}
	return http.StatusOK, user, nil
}
