//go:build integration

package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo/transactor"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FlatsTestSuite struct {
	suite.Suite
	ctx        context.Context
	flatRepo   repo.Flat
	houseRepo  repo.House
	transactor *transactor.TransactionManager
}

func TestFlats(t *testing.T) {
	suite.Run(t, new(FlatsTestSuite))
}

func (s *FlatsTestSuite) SetupSuite() {
	s.T().Parallel()
	s.transactor = transactor.NewTransactionManager(db.GetPool())
	s.flatRepo = repo.NewFlatRepo(s.transactor)
	s.houseRepo = repo.NewHouseRepo(s.transactor)
	s.ctx = context.Background()
}

func (s *FlatsTestSuite) SetupTest() {
	s.T().Parallel()
}

func (s *FlatsTestSuite) TestCreate() {
	_ = s.createFlat()
}

func (s *FlatsTestSuite) TestGetByID() {
	flat := s.createFlat()

	actual, err := s.flatRepo.GetByID(s.ctx, flat.ID)
	require.NoError(s.T(), err)
	require.EqualExportedValues(s.T(), flat, actual)
}

func (s *FlatsTestSuite) TestUpdateStatus() {
	flat := s.createFlat()
	flat.Status = domain.FlatApproved

	err := s.flatRepo.UpdateStatus(s.ctx, flat.ID, flat.Status)
	require.NoError(s.T(), err)

	actual, err := s.flatRepo.GetByID(s.ctx, flat.ID)
	require.NoError(s.T(), err)
	require.EqualExportedValues(s.T(), flat, actual)
}

func (s *FlatsTestSuite) TestSetModeratorID() {
	flat := s.createFlat()
	flat.ModeratorID = uuid.New()

	err := s.flatRepo.SetModeratorID(s.ctx, flat.ID, &flat.ModeratorID)
	require.NoError(s.T(), err)

	actual, err := s.flatRepo.GetByID(s.ctx, flat.ID)
	require.NoError(s.T(), err)
	require.EqualValues(s.T(), flat.ModeratorID, actual.ModeratorID)
}

func (s *FlatsTestSuite) createFlat() domain.Flat {
	houses := NewHouse()

	houseID, err := s.houseRepo.Create(s.ctx, houses)
	require.NoError(s.T(), err)

	houses.ID = houseID
	flat := NewFlats(houseID)

	res, err := s.flatRepo.Create(s.ctx, flat)
	require.NoError(s.T(), err)

	flat.ID = res.ID
	flat.Number = res.Number
	return flat
}
