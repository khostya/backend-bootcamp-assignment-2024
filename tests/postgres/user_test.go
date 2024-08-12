//go:build integration

package postgres

import (
	"backend-bootcamp-assignment-2024/internal/repo"
	"backend-bootcamp-assignment-2024/internal/repo/transactor"
	"context"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UsersTestSuite struct {
	suite.Suite
	ctx        context.Context
	userRepo   repo.User
	transactor *transactor.TransactionManager
}

func TestUsers(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}

func (s *UsersTestSuite) SetupSuite() {
	s.T().Parallel()
	s.transactor = transactor.NewTransactionManager(db.GetPool())
	s.userRepo = repo.NewUserRepo(s.transactor)
	s.ctx = context.Background()
}

func (s *UsersTestSuite) TestCreate() {
	user := NewUser()

	err := s.userRepo.Create(s.ctx, user)
	require.NoError(s.T(), err)
}

func (s *UsersTestSuite) TestGetByID() {
	user := NewUser()

	err := s.userRepo.Create(s.ctx, user)
	require.NoError(s.T(), err)

	actual, err := s.userRepo.GetByID(s.ctx, user.ID)
	require.NoError(s.T(), err)
	require.EqualExportedValues(s.T(), user, actual)
}
