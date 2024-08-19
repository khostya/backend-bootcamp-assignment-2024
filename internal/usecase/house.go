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
		GetFullByID(ctx context.Context, id uint, flatStatus *domain.FlatStatus) (domain.House, error)
	}

	subscriptionRepo interface {
		Create(ctx context.Context, subscription domain.Subscription) error
		GetByHouseID(ctx context.Context, houseID uint) ([]domain.Subscription, error)
	}

	House struct {
		transactionManager transactionManager
		houseRepo          houseRepo
		subscriptionRepo   subscriptionRepo
	}
)

func NewHouseUseCase(repo houseRepo, subscriptionRepo subscriptionRepo, manager transactionManager) House {
	return House{
		houseRepo:          repo,
		transactionManager: manager,
		subscriptionRepo:   subscriptionRepo,
	}
}

func (uc House) GetByID(ctx context.Context, id uint, userType domain.UserType) (domain.House, error) {
	if userType == domain.UserClient {
		status := domain.FlatApproved
		return uc.houseRepo.GetFullByID(ctx, id, &status)
	}

	return uc.houseRepo.GetFullByID(ctx, id, nil)
}

func (uc House) Create(ctx context.Context, param dto.CreateHouseParam) (domain.House, error) {
	house := domain.NewHouse(param)

	id, err := uc.houseRepo.Create(ctx, house)
	house.ID = id

	return house, err
}

func (uc House) Subscribe(ctx context.Context, subscription domain.Subscription) error {
	return uc.subscriptionRepo.Create(ctx, subscription)
}
