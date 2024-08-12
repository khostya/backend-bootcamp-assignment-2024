package http

import (
	"backend-bootcamp-assignment-2024/internal/cache"
	"backend-bootcamp-assignment-2024/internal/domain"
	"backend-bootcamp-assignment-2024/internal/http/api"
	mock_usecase "backend-bootcamp-assignment-2024/internal/usecase/mocks"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"testing"
)

type mocks struct {
	House    *mock_usecase.MockhouseService
	Flat     *mock_usecase.MockflatService
	Auth     *mock_usecase.MockauthService
	useCases UseCases
}

func newMocks(t *testing.T) mocks {
	ctrl := gomock.NewController(t)
	flat := mock_usecase.NewMockflatService(ctrl)
	house := mock_usecase.NewMockhouseService(ctrl)
	auth := mock_usecase.NewMockauthService(ctrl)

	return mocks{
		Flat:  flat,
		House: house,
		Auth:  auth,
		useCases: UseCases{
			flat, house, auth,
		},
	}
}

func TestServerAuth_Login(t *testing.T) {
	t.Parallel()

	type test struct {
		name    string
		input   api.PostLoginJSONBody
		mockFn  func(m mocks)
		status  int
		wantErr bool
	}

	tests := []test{
		{
			name: "ok",
			input: api.PostLoginJSONBody{
				Id:       uuid.New(),
				Password: uuid.NewString(),
			},
			wantErr: false,
			status:  http.StatusOK,
			mockFn: func(m mocks) {
				m.Auth.EXPECT().Login(gomock.Any(), gomock.Any()).Times(1).Return(domain.Token("#1"), nil)
			},
		},
		{
			name: "bad request_password required",
			input: api.PostLoginJSONBody{
				Id: uuid.New(),
			},
			wantErr: true,
			status:  http.StatusBadRequest,
			mockFn: func(m mocks) {
			},
		},
		{
			name: "bad request_id required",
			input: api.PostLoginJSONBody{
				Password: uuid.NewString(),
			},
			wantErr: true,
			status:  http.StatusBadRequest,
			mockFn: func(m mocks) {
			},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mocks := newMocks(t)
			tt.mockFn(mocks)

			server, err := newServer(mocks.useCases, Cache{house: cache.NewHouseCache(0, 0)})
			require.NoError(t, err)

			status, _, err := server.postLogin(ctx, tt.input)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.status, status)
		})
	}
}

func TestServerAuth_Register(t *testing.T) {
	t.Parallel()

	type test struct {
		name    string
		input   api.PostRegisterJSONBody
		mockFn  func(m mocks)
		status  int
		wantErr bool
	}

	email := api.Email("fff@gmail.com")
	tests := []test{
		{
			name: "ok without email",
			input: api.PostRegisterJSONBody{
				Email:    nil,
				Password: uuid.NewString(),
				UserType: api.Client,
			},
			wantErr: false,
			status:  http.StatusOK,
			mockFn: func(m mocks) {
				m.Auth.EXPECT().Register(gomock.Any(), gomock.Any()).Times(1).Return(uuid.UUID{}, nil)
			},
		},
		{
			name: "ok with email",
			input: api.PostRegisterJSONBody{
				Email:    &email,
				Password: uuid.NewString(),
				UserType: api.Client,
			},
			wantErr: false,
			status:  http.StatusOK,
			mockFn: func(m mocks) {
				m.Auth.EXPECT().Register(gomock.Any(), gomock.Any()).Times(1).Return(uuid.UUID{}, nil)
			},
		},
		{
			name: "bad request_password required",
			input: api.PostRegisterJSONBody{
				UserType: api.Client,
			},
			wantErr: true,
			status:  http.StatusBadRequest,
			mockFn:  func(m mocks) {},
		},
		{
			name: "bad request_password required",
			input: api.PostRegisterJSONBody{
				Password: uuid.NewString(),
			},
			wantErr: true,
			status:  http.StatusBadRequest,
			mockFn:  func(m mocks) {},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mocks := newMocks(t)
			tt.mockFn(mocks)

			server, err := newServer(mocks.useCases, Cache{house: cache.NewHouseCache(0, 0)})
			require.NoError(t, err)

			status, _, err := server.postRegister(ctx, tt.input)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.status, status)
		})
	}
}
