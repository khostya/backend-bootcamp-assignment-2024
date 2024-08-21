package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/dto"
	mock_repo "github.com/khostya/backend-bootcamp-assignment-2024/internal/repo/mocks"
	mock_transactor "github.com/khostya/backend-bootcamp-assignment-2024/internal/repo/transactor/mocks"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/auth"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/hash"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

type authMocks struct {
	mockUserRepo   *mock_repo.MockuserRepo
	mockTransactor *mock_transactor.MockTransactor
}

func newAuthMocks(t *testing.T) authMocks {
	ctrl := gomock.NewController(t)

	return authMocks{
		mockTransactor: mock_transactor.NewMockTransactor(ctrl),
		mockUserRepo:   mock_repo.NewMockuserRepo(ctrl),
	}
}

func TestAuthUseCase_DummyLogin(t *testing.T) {
	t.Parallel()

	type test struct {
		name     string
		userType domain.UserType
	}

	ctx := context.Background()
	tests := []test{
		{
			name:     "ok",
			userType: domain.UserClient,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mocks := newAuthMocks(t)

			authUseCase := NewAuthUseCase(AuthDeps{
				TransactionManager: mocks.mockTransactor,
				UserRepo:           mocks.mockUserRepo,
				PasswordHasher:     hash.NewPasswordHasher(2),
				TokenManager:       auth.NewManager("312312"),
				AccessTokenTTL:     time.Hour,
			})

			token, err := authUseCase.DummyLogin(ctx, tt.userType)
			require.NoError(t, err)
			require.NotEqual(t, "", string(token))
		})
	}
}

func TestAuthUseCase_Login(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()
	)

	type test struct {
		name            string
		input           dto.LoginUserParam
		currentPassword string
		err             error
		mockFn          func(ctx context.Context, param dto.LoginUserParam, hashedPassword string, m authMocks)
	}

	tests := []test{
		{
			name: "incorrect password",
			input: dto.LoginUserParam{
				Id:       uuid.New(),
				Password: "31231",
			},
			currentPassword: "3",
			mockFn: func(ctx context.Context, param dto.LoginUserParam, hashedPassword string, m authMocks) {
				m.mockUserRepo.EXPECT().GetByID(ctx, param.Id).
					Times(1).Return(domain.User{Password: hashedPassword}, nil)
			},
			err: ErrIncorrectPassword,
		},
		{
			name: "ok",
			input: dto.LoginUserParam{
				Id:       uuid.New(),
				Password: "3",
			},
			currentPassword: "3",
			mockFn: func(ctx context.Context, param dto.LoginUserParam, hashedPassword string, m authMocks) {
				m.mockUserRepo.EXPECT().GetByID(ctx, param.Id).
					Times(1).Return(domain.User{Password: hashedPassword}, nil)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			hasher := hash.NewPasswordHasher(2)
			hashedPassword, err := hasher.Hash(tt.currentPassword)
			require.NoError(t, err)

			mocks := newAuthMocks(t)
			tt.mockFn(ctx, tt.input, hashedPassword, mocks)

			authUseCase := NewAuthUseCase(AuthDeps{
				TransactionManager: mocks.mockTransactor,
				UserRepo:           mocks.mockUserRepo,
				PasswordHasher:     hasher,
				TokenManager:       auth.NewManager("312312"),
				AccessTokenTTL:     time.Hour,
			})

			_, err = authUseCase.Login(ctx, tt.input)
			require.ErrorIs(t, err, tt.err)
		})
	}
}

func TestAuthUseCase_Register(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()
	)

	type test struct {
		name            string
		input           dto.RegisterUserParam
		currentPassword string
		mockFn          func(ctx context.Context, param dto.RegisterUserParam, m authMocks)
	}

	tests := []test{
		{
			name: "ok",
			input: dto.RegisterUserParam{
				Email:    "3123@gmail.com",
				Password: "#!@31",
				UserType: string(domain.UserClient),
			},
			mockFn: func(ctx context.Context, param dto.RegisterUserParam, m authMocks) {
				m.mockUserRepo.EXPECT().Create(ctx, gomock.Any()).
					Times(1).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mocks := newAuthMocks(t)
			tt.mockFn(ctx, tt.input, mocks)

			authUseCase := NewAuthUseCase(AuthDeps{
				TransactionManager: mocks.mockTransactor,
				UserRepo:           mocks.mockUserRepo,
				PasswordHasher:     hash.NewPasswordHasher(2),
				TokenManager:       auth.NewManager("312312"),
				AccessTokenTTL:     time.Hour,
			})

			_, err := authUseCase.Register(ctx, tt.input)
			require.NoError(t, err)
		})
	}
}
