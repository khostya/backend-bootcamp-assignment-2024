package usecase

import (
	"context"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/dto"
	"time"
)

type (
	houseRepo interface {
		GetByID(ctx context.Context, id uint) (domain.House, error)
		Create(ctx context.Context, house domain.House) (uint, error)
		UpdateLastFlatAddedAt(ctx context.Context, id uint, updatedAt time.Time) error
	}

	House struct {
		transactionManager transactionManager
		houseRepo          houseRepo
	}
)

func NewHouseUseCase(repo houseRepo, manager transactionManager) House {
	return House{
		houseRepo:          repo,
		transactionManager: manager,
	}
}

func (uc House) GetByID(ctx context.Context, id uint) (domain.House, error) {
	return uc.houseRepo.GetByID(ctx, id)
}

func (uc House) Create(ctx context.Context, param dto.CreateHouseParam) (domain.House, error) {
	house := domain.NewHouse(param)

	id, err := uc.houseRepo.Create(ctx, house)
	house.ID = id

	return house, err
}

func (uc House) Subscribe(ctx context.Context, id int, email string) error {
	panic("11")
}
