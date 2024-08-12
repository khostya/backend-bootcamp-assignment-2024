package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/dto"
	"time"
)

type (
	flatRepo interface {
		Create(ctx context.Context, flat domain.Flat) (uint, error)
		UpdateStatus(ctx context.Context, id uint, status domain.FlatStatus) error
		GetByID(ctx context.Context, id uint) (domain.Flat, error)
		SetModeratorID(ctx context.Context, id uint, moderatorID *uuid.UUID) error
	}

	Flat struct {
		transactionManager transactionManager
		flatRepo           flatRepo
		houseRepo          houseRepo
	}
)

func NewFlatUseCase(repo flatRepo, houseRepo houseRepo, manager transactionManager) Flat {
	return Flat{
		houseRepo:          houseRepo,
		flatRepo:           repo,
		transactionManager: manager,
	}
}

func (uc Flat) Create(ctx context.Context, param dto.CreateFlatParam) (domain.Flat, error) {
	flat := domain.NewFlat(param)

	err := uc.transactionManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		ID, err := uc.flatRepo.Create(ctx, flat)
		if err != nil {
			return err
		}

		flat.ID = ID
		return uc.houseRepo.UpdateLastFlatAddedAt(ctx, flat.HouseID, time.Now())
	})

	return flat, uc.transactionManager.Unwrap(err)
}

func (uc Flat) Update(ctx context.Context, param dto.UpdateFlatParam) (domain.Flat, error) {
	var (
		err  error
		flat domain.Flat
	)

	err = uc.transactionManager.RunRepeatableRead(ctx, func(ctx context.Context) error {
		flat, err = uc.flatRepo.GetByID(ctx, param.Id)
		if err != nil {
			return err
		}

		if domain.FlatStatus(param.Status) == domain.FlatModeration &&
			flat.Status == domain.FlatCreated {
			err := uc.flatRepo.SetModeratorID(ctx, param.Id, nil)
			if err != nil {
				return err
			}
		}

		if flat.Status == domain.FlatModeration && flat.ModeratorID == param.ModeratorID {
			err := uc.flatRepo.SetModeratorID(ctx, param.Id, nil)
			if err != nil {
				return err
			}
		}

		flat.Status = domain.FlatStatus(param.Status)
		return uc.flatRepo.UpdateStatus(ctx, param.Id, domain.FlatStatus(param.Status))
	})

	if err != nil {
		return domain.Flat{}, uc.transactionManager.Unwrap(err)
	}

	return flat, nil
}
