// Code generated by ifacemaker. DO NOT EDIT.

package mock_repo

import (
	"context"

	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
)

// subscriptionRepo ...
type subscriptionRepo interface {
	Create(ctx context.Context, subscription domain.Subscription) error
	GetByHouseID(ctx context.Context, houseID uint) ([]domain.Subscription, error)
}
