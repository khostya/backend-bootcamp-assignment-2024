package usecase

import (
	"backend-bootcamp-assignment-2024/internal/domain"
	"backend-bootcamp-assignment-2024/internal/dto"
	mock_repo "backend-bootcamp-assignment-2024/internal/repo/mocks"
	mock_transactor "backend-bootcamp-assignment-2024/internal/repo/transactor/mocks"
	"context"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

type houseMocks struct {
	mockHouseRepo  *mock_repo.MockhouseRepo
	mockTransactor *mock_transactor.MockTransactor
}

func newHouseMocks(t *testing.T) houseMocks {
	ctrl := gomock.NewController(t)

	return houseMocks{
		mockTransactor: mock_transactor.NewMockTransactor(ctrl),
		mockHouseRepo:  mock_repo.NewMockhouseRepo(ctrl),
	}
}

func TestHouseUseCase_Create(t *testing.T) {
	t.Parallel()

	type test struct {
		name   string
		input  dto.CreateHouseParam
		mockFn func(ctx context.Context, m houseMocks)
	}

	ctx := context.Background()
	tests := []test{
		{
			name: "ok",
			input: dto.CreateHouseParam{
				Address:   "#131",
				Developer: "31312",
				Year:      314124,
			},
			mockFn: func(ctx context.Context, m houseMocks) {
				m.mockHouseRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(uint(1), nil)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mocks := newHouseMocks(t)
			tt.mockFn(ctx, mocks)

			houseUseCase := NewHouseUseCase(mocks.mockHouseRepo, mocks.mockTransactor)

			_, err := houseUseCase.Create(ctx, tt.input)
			require.NoError(t, err)
		})
	}
}

func TestHouseUseCase_GetByID(t *testing.T) {
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
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mocks := newFlatMocks(t)
			tt.mockFn(ctx, tt.input, mocks)

			flatUseCase := NewFlatUseCase(mocks.mockFlatRepo, mocks.mockTransactor)

			_, err := flatUseCase.Update(ctx, tt.input)
			require.ErrorIs(t, err, tt.err)
		})
	}
}
