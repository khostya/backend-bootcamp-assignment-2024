package http

import (
	"context"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/cache"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/api"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/middleware"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"testing"
)

func TestServerHouse_postCreate(t *testing.T) {
	t.Parallel()

	type test struct {
		ctx     context.Context
		name    string
		input   api.PostHouseCreateJSONBody
		mockFn  func(ctx context.Context, m mocks)
		status  int
		wandErr bool
	}

	ctx := context.Background()

	developer := api.Developer("31231")
	tests := []test{
		{
			ctx:  context.WithValue(ctx, middleware.KeyUserType, domain.UserModerator),
			name: "ok without developer",
			input: api.PostHouseCreateJSONBody{
				Address: "4141",
				Year:    2024,
			},
			wandErr: false,
			status:  http.StatusOK,
			mockFn: func(ctx context.Context, m mocks) {
				m.House.EXPECT().Create(ctx, gomock.Any()).Times(1).Return(domain.House{}, nil)
			},
		},
		{
			ctx:  context.WithValue(ctx, middleware.KeyUserType, domain.UserModerator),
			name: "ok",
			input: api.PostHouseCreateJSONBody{
				Developer: &developer,
				Address:   "4141",
				Year:      2024,
			},
			wandErr: false,
			status:  http.StatusOK,
			mockFn: func(ctx context.Context, m mocks) {
				m.House.EXPECT().Create(ctx, gomock.Any()).Times(1).Return(domain.House{}, nil)
			},
		},
		{
			ctx:  context.WithValue(ctx, middleware.KeyUserType, domain.UserModerator),
			name: "bad request without address",
			input: api.PostHouseCreateJSONBody{
				Year: 2024,
			},
			wandErr: true,
			status:  http.StatusBadRequest,
			mockFn:  func(ctx context.Context, m mocks) {},
		},
		{
			ctx:  context.WithValue(ctx, middleware.KeyUserType, domain.UserModerator),
			name: "bad request without year",
			input: api.PostHouseCreateJSONBody{
				Address: "#1231",
			},
			wandErr: true,
			status:  http.StatusBadRequest,
			mockFn:  func(ctx context.Context, m mocks) {},
		},
		{
			ctx:  context.WithValue(ctx, middleware.KeyUserType, domain.UserClient),
			name: "unauthorized with client",
			input: api.PostHouseCreateJSONBody{
				Address: "#1231",
			},
			wandErr: true,
			status:  http.StatusUnauthorized,
			mockFn:  func(ctx context.Context, m mocks) {},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mocks := newMocks(t)
			tt.mockFn(tt.ctx, mocks)

			server, err := newServer(mocks.useCases, Cache{House: cache.NewHouseCache(0, 0)})
			require.NoError(t, err)

			status, _, err := server.postHouseCreate(tt.ctx, tt.input)
			require.Equal(t, tt.wandErr, err != nil)
			require.Equal(t, tt.status, status)
		})
	}
}
