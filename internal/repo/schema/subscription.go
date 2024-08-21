package schema

import (
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
)

type (
	Subscription struct {
		UserEmail string `db:"user_email"`
		HouseID   uint   `db:"house_id"`
	}
)

func (s Subscription) Values() []any {
	return []any{s.HouseID, s.UserEmail}
}

func (s Subscription) Columns() []string {
	return []string{"house_id", "user_email"}
}

func NewSubscription(subscription domain.Subscription) Subscription {
	return Subscription{
		HouseID:   subscription.HouseID,
		UserEmail: subscription.UserEmail,
	}
}

func NewDomainSubscription(subscription Subscription) domain.Subscription {
	return domain.Subscription{
		UserEmail: subscription.UserEmail,
		HouseID:   subscription.HouseID,
	}
}

func NewDomainSubscriptions(subscriptions []Subscription) []domain.Subscription {
	var res []domain.Subscription

	for _, s := range subscriptions {
		res = append(res, NewDomainSubscription(s))
	}

	return res
}
