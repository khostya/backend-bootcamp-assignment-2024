//go:build integration

package postgres

import (
	"context"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo/transactor"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"sort"
	"testing"
)

type SubscriptionTestSuite struct {
	suite.Suite
	ctx              context.Context
	transactor       *transactor.TransactionManager
	userRepo         repo.User
	houseRepo        repo.House
	subscriptionRepo repo.Subscription
}

func TestSubscription(t *testing.T) {
	suite.Run(t, new(SubscriptionTestSuite))
}

func (s *SubscriptionTestSuite) SetupSuite() {
	s.T().Parallel()
	s.transactor = transactor.NewTransactionManager(db.GetPool())
	s.houseRepo = repo.NewHouseRepo(s.transactor)
	s.userRepo = repo.NewUserRepo(s.transactor)
	s.subscriptionRepo = repo.NewSubscriptionRepo(s.transactor)
	s.ctx = context.Background()
}

func (s *SubscriptionTestSuite) SetupTest() {
	s.T().Parallel()
}

func (s *SubscriptionTestSuite) TestCreate() {
	_ = s.create()
}

func (s *SubscriptionTestSuite) TestGetByID() {
	subscription1 := s.create()

	user := NewUser()
	err := s.userRepo.Create(s.ctx, user)
	require.NoError(s.T(), err)

	subscription2 := domain.Subscription{HouseID: subscription1.HouseID, UserEmail: user.Email}
	err = s.subscriptionRepo.Create(s.ctx, subscription2)
	require.NoError(s.T(), err)

	subscriptions, err := s.subscriptionRepo.GetByHouseID(s.ctx, subscription1.HouseID)
	require.NoError(s.T(), err)

	sort.Slice(subscriptions, func(i, j int) bool {
		return subscriptions[i].UserEmail < subscriptions[j].UserEmail
	})

	expected := []domain.Subscription{subscription1, subscription2}
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].UserEmail < expected[j].UserEmail
	})

	require.Equal(s.T(), expected, subscriptions)
}

func (s *SubscriptionTestSuite) create() domain.Subscription {
	houses := NewHouse()

	houseID, err := s.houseRepo.Create(s.ctx, houses)
	require.NoError(s.T(), err)

	houses.ID = houseID

	user := NewUser()
	err = s.userRepo.Create(s.ctx, user)
	require.NoError(s.T(), err)

	subscription := domain.Subscription{HouseID: houseID, UserEmail: user.Email}

	err = s.subscriptionRepo.Create(s.ctx, subscription)
	require.NoError(s.T(), err)

	return subscription
}
