package usecase

import (
	"context"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/dto"
	mock_repo "github.com/khostya/backend-bootcamp-assignment-2024/internal/repo/mocks"
	mock_transactor "github.com/khostya/backend-bootcamp-assignment-2024/internal/repo/transactor/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

type flatMocks struct {
	mockFlatRepo   *mock_repo.MockflatRepo
	mockHouseRepo  *mock_repo.MockhouseRepo
	mockTransactor *mock_transactor.MockTransactor
}

func newFlatMocks(t *testing.T) flatMocks {
	ctrl := gomock.NewController(t)

	return flatMocks{
		mockHouseRepo:  mock_repo.NewMockhouseRepo(ctrl),
		mockTransactor: mock_transactor.NewMockTransactor(ctrl),
		mockFlatRepo:   mock_repo.NewMockflatRepo(ctrl),
	}
}

func TestFlatUseCase_Create(t *testing.T) {
	t.Parallel()

	type test struct {
		name   string
		input  dto.CreateFlatParam
		mockFn func(ctx context.Context, m flatMocks)
	}

	ctx := context.Background()
	tests := []test{
		{
			name: "ok",
			input: dto.CreateFlatParam{
				HouseID: 1,
				Price:   1,
				Rooms:   1,
			},

			mockFn: func(ctx context.Context, m flatMocks) {
				m.mockFlatRepo.EXPECT().Create(gomock.Any(), gomock.Any()).
					Times(1).Return(uint(1), nil)
				m.mockHouseRepo.EXPECT().UpdateLastFlatAddedAt(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).Return(nil)
				m.mockTransactor.EXPECT().RunRepeatableRead(gomock.Any(), gomock.Any()).Times(1).Return(nil).
					DoAndReturn(func(ctx context.Context, transaction func(ctx context.Context) error) error {
						return transaction(ctx)
					})
				m.mockTransactor.EXPECT().Unwrap(gomock.Any()).Times(1).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mocks := newFlatMocks(t)
			tt.mockFn(ctx, mocks)

			flatUseCase := NewFlatUseCase(mocks.mockFlatRepo, mocks.mockHouseRepo, mocks.mockTransactor)

			_, err := flatUseCase.Create(ctx, tt.input)
			require.NoError(t, err)
		})
	}
}

func TestFlatUseCase_Update(t *testing.T) {
	t.Parallel()

	type test struct {
		name   string
		input  dto.UpdateFlatParam
		mockFn func(ctx context.Context, param dto.UpdateFlatParam, m flatMocks)
		err    error
	}

	ctx := context.Background()
	tests := []test{
		{
			name: "ok",
			input: dto.UpdateFlatParam{
				Id:     1,
				Status: string(domain.FlatApproved),
			},
			mockFn: func(ctx context.Context, param dto.UpdateFlatParam, m flatMocks) {
				m.mockFlatRepo.EXPECT().UpdateStatus(ctx, param.Id, domain.FlatStatus(param.Status)).
					Times(1).Return(nil)
				m.mockFlatRepo.EXPECT().GetByID(ctx, param.Id).
					Times(1).Return(domain.Flat{}, nil)
				m.mockTransactor.EXPECT().RunRepeatableRead(gomock.Any(), gomock.Any()).Times(1).Return(nil).
					DoAndReturn(func(ctx context.Context, transaction func(ctx context.Context) error) error {
						return transaction(ctx)
					})
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mocks := newFlatMocks(t)
			tt.mockFn(ctx, tt.input, mocks)

			flatUseCase := NewFlatUseCase(mocks.mockFlatRepo, mocks.mockHouseRepo, mocks.mockTransactor)

			_, err := flatUseCase.Update(ctx, tt.input)
			require.ErrorIs(t, err, tt.err)
		})
	}
}
