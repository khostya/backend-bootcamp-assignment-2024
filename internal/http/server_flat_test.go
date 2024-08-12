package http

import (
	"backend-bootcamp-assignment-2024/internal/cache"
	"backend-bootcamp-assignment-2024/internal/domain"
	"backend-bootcamp-assignment-2024/internal/http/api"
	"backend-bootcamp-assignment-2024/internal/http/middleware"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"testing"
)

func TestServerFlat_postCreate(t *testing.T) {
	t.Parallel()

	type test struct {
		name    string
		input   api.PostFlatCreateJSONBody
		mockFn  func(m mocks)
		status  int
		wantErr bool
	}

	tests := []test{
		{
			name: "ok",
			input: api.PostFlatCreateJSONBody{
				HouseId: 1312,
				Price:   3123,
				Rooms:   31312,
			},
			wantErr: false,
			status:  http.StatusOK,
			mockFn: func(m mocks) {
				m.Flat.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(domain.Flat{}, nil)
			},
		},
		{
			name: "bad request without house id",
			input: api.PostFlatCreateJSONBody{
				Price: 3123,
				Rooms: 31312,
			},
			wantErr: true,
			status:  http.StatusBadRequest,
			mockFn: func(m mocks) {
			},
		},
		{
			name: "bad request without price",
			input: api.PostFlatCreateJSONBody{
				HouseId: 3131,
				Rooms:   31312,
			},
			wantErr: true,
			status:  http.StatusBadRequest,
			mockFn: func(m mocks) {
			},
		},
		{
			name: "bad request without rooms",
			input: api.PostFlatCreateJSONBody{
				Price:   31312,
				HouseId: 3131,
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
			t.Parallel()

			mocks := newMocks(t)
			tt.mockFn(mocks)

			server, err := newServer(mocks.useCases, Cache{House: cache.NewHouseCache(0, 0)})
			require.NoError(t, err)

			status, _, err := server.postFlatCreate(ctx, tt.input)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.status, status)
		})
	}
}

func TestServerFlat_postUpdate(t *testing.T) {
	t.Parallel()

	type test struct {
		name    string
		input   api.PostFlatUpdateJSONBody
		mockFn  func(m mocks)
		status  int
		wandErr bool
	}

	tests := []test{
		{
			name: "ok",
			input: api.PostFlatUpdateJSONBody{
				Id:     1,
				Status: api.Created,
			},
			wandErr: false,
			status:  http.StatusOK,
			mockFn: func(m mocks) {
				m.Flat.EXPECT().Update(gomock.Any(), gomock.Any()).Times(1).Return(domain.Flat{}, nil)
			},
		},
		{
			name: "bad request without status",
			input: api.PostFlatUpdateJSONBody{
				Id: 1,
			},
			wandErr: true,
			status:  http.StatusBadRequest,
			mockFn: func(m mocks) {
			},
		},
		{
			name: "bad request without id",
			input: api.PostFlatUpdateJSONBody{
				Status: api.Created,
			},
			wandErr: true,
			status:  http.StatusBadRequest,
			mockFn: func(m mocks) {
			},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mocks := newMocks(t)
			tt.mockFn(mocks)

			server, err := newServer(mocks.useCases, Cache{House: cache.NewHouseCache(0, 0)})
			require.NoError(t, err)

			ctx := context.WithValue(ctx, middleware.UserID, uuid.New())
			ctx = context.WithValue(ctx, middleware.UserType, domain.UserModerator)

			status, _, err := server.postFlatUpdate(ctx, tt.input)
			require.Equal(t, tt.wandErr, err != nil)
			require.Equal(t, tt.status, status)
		})
	}
}
