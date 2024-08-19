package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo/exec"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo/schema"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo/transactor"
)

const (
	subscriptionTable = "bootcamp.subscriptions"
)

type (
	Subscription struct {
		queryEngineProvider transactor.QueryEngineProvider
	}
)

func NewSubscriptionRepo(provider transactor.QueryEngineProvider) Subscription {
	return Subscription{provider}
}

func (s Subscription) Create(ctx context.Context, subscription domain.Subscription) error {
	db := s.queryEngineProvider.GetQueryEngine(ctx)

	record := schema.NewSubscription(subscription)
	query := sq.Insert(subscriptionTable).
		Columns(record.Columns()...).
		Values(record.Values()...).
		PlaceholderFormat(sq.Dollar)

	return exec.Insert(ctx, query, db)
}

func (s Subscription) GetByHouseID(ctx context.Context, houseID uint) ([]domain.Subscription, error) {
	db := s.queryEngineProvider.GetQueryEngine(ctx)

	query := sq.Select(schema.Subscription{}.Columns()...).
		From(subscriptionTable).
		Where("house_id = $1", houseID).
		PlaceholderFormat(sq.Dollar)

	all, err := exec.ScanALL[schema.Subscription](ctx, query, db)
	if err != nil {
		return nil, err
	}

	return schema.NewDomainSubscriptions(all), nil
}
