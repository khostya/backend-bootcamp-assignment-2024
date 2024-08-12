package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/dto"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/auth"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/hash"
	"time"
)

type (
	userRepo interface {
		GetByID(ctx context.Context, id uuid.UUID) (domain.User, error)
		Create(ctx context.Context, user domain.User) error
	}

	Auth struct {
		transactionManager transactionManager
		userRepo           userRepo

		passwordHasher hash.PasswordHasher
		tokenManager   auth.TokenManager

		accessTokenTTL time.Duration
	}

	AuthDeps struct {
		TransactionManager transactionManager
		UserRepo           userRepo

		PasswordHasher hash.PasswordHasher
		TokenManager   auth.TokenManager

		AccessTokenTTL time.Duration
	}
)

func NewAuthUseCase(deps AuthDeps) Auth {
	return Auth{
		userRepo:           deps.UserRepo,
		transactionManager: deps.TransactionManager,
		passwordHasher:     deps.PasswordHasher,
		tokenManager:       deps.TokenManager,
		accessTokenTTL:     deps.AccessTokenTTL,
	}
}

func (uc Auth) DummyLogin(_ context.Context, userType domain.UserType) (domain.Token, error) {
	return uc.createToken(uuid.New(), userType)
}

func (uc Auth) Login(ctx context.Context, param dto.LoginUserParam) (domain.Token, error) {
	user, err := uc.userRepo.GetByID(ctx, param.Id)
	if err != nil {
		return "", err
	}

	if !uc.passwordHasher.Equals(user.Password, param.Password) {
		return "", ErrIncorrectPassword
	}

	return uc.createToken(user.ID, user.UserType)
}

func (uc Auth) Register(ctx context.Context, param dto.RegisterUserParam) (uuid.UUID, error) {
	hashedPassword, err := uc.passwordHasher.Hash(param.Password)
	if err != nil {
		return uuid.UUID{}, err
	}

	user := domain.NewUser(param, hashedPassword)

	err = uc.userRepo.Create(ctx, user)

	return user.ID, err
}

func (uc Auth) createToken(ID uuid.UUID, userType domain.UserType) (domain.Token, error) {
	token, err := uc.tokenManager.NewUserJWT(ID, string(userType), time.Now().Add(uc.accessTokenTTL))
	return domain.Token(token), err
}
