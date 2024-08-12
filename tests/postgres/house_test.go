//go:build integration

package postgres

import (
	"context"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo/transactor"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HousesTestSuite struct {
	suite.Suite
	ctx        context.Context
	houseRepo  repo.House
	transactor *transactor.TransactionManager
}

func TestHouses(t *testing.T) {
	suite.Run(t, new(HousesTestSuite))
}

func (s *HousesTestSuite) SetupSuite() {
	s.T().Parallel()
	s.transactor = transactor.NewTransactionManager(db.GetPool())
	s.houseRepo = repo.NewHouseRepo(s.transactor)
	s.ctx = context.Background()
}

func (s *HousesTestSuite) TestCreate() {
	houses := NewHouses()

	_, err := s.houseRepo.Create(s.ctx, houses)
	require.NoError(s.T(), err)
}

func (s *HousesTestSuite) TestGetByID() {
	houses := NewHouses()

	id, err := s.houseRepo.Create(s.ctx, houses)
	require.NoError(s.T(), err)
	houses.ID = id

	actual, err := s.houseRepo.GetByID(s.ctx, id)
	require.NoError(s.T(), err)
	require.EqualExportedValues(s.T(), houses, actual)
}
