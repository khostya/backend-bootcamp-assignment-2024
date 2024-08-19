// Code generated by ifacemaker. DO NOT EDIT.

package mock_repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/dto"
)

// flatRepo ...
type flatRepo interface {
	Create(ctx context.Context, flat domain.Flat) (dto.FlatCreateResult, error)
	GetByID(ctx context.Context, id uint) (domain.Flat, error)
	UpdateStatus(ctx context.Context, id uint, status domain.FlatStatus) error
	SetModeratorID(ctx context.Context, id uint, moderatorID *uuid.UUID) error
}
