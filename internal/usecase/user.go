package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
)

type (
	User struct {
		userRepo userRepo
	}
)

func NewUserUseCase(userRepo userRepo) User {
	return User{
		userRepo: userRepo,
	}
}

func (uc User) GetByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return uc.userRepo.GetByID(ctx, id)
}
