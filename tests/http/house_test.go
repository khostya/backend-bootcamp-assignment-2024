//go:build integration

package http

import (
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/api"
	model "github.com/khostya/backend-bootcamp-assignment-2024/internal/http/api/models"
	"github.com/khostya/backend-bootcamp-assignment-2024/tests/http/client"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type (
	HouseTestSuite struct {
		suite.Suite
		ctx    context.Context
		client *api.Client
	}
)

func TestHouse(t *testing.T) {
	suite.Run(t, new(HouseTestSuite))
}

func (s *HouseTestSuite) SetupSuite() {
	s.T().Parallel()
	s.client = client.New()
	s.ctx = context.Background()
}

func (s *HouseTestSuite) SetupTest() {
	s.T().Parallel()
}

func (s *HouseTestSuite) TestCreate() {
	createHouse(s.ctx, s.T())
}

func createHouse(ctx context.Context, t *testing.T) (*api.House, userResult, api.Token) {
	client := client.New()

	user, token := login(ctx, t, model.Moderator)

	resp, err := client.PostHouseCreate(ctx, api.PostHouseCreateJSONRequestBody{
		Address:   gofakeit.Address().Address,
		Developer: nil,
		Year:      312321,
	}, addAuth(token))
	require.NoError(t, err)

	response, err := api.ParsePostHouseCreateResponse(resp)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())

	return response.JSON200, user, token
}

func (s *HouseTestSuite) TestSubscribe() {
	house, user, token := createHouse(s.ctx, s.T())

	resp, err := s.client.PostHouseIdSubscribe(s.ctx, house.Id, api.PostHouseIdSubscribeJSONRequestBody{
		Email: user.Email,
	}, addAuth(token))
	require.NoError(s.T(), err)

	response, err := api.ParsePostHouseIdSubscribeResponse(resp)
	require.NoError(s.T(), err)
	require.Equal(s.T(), http.StatusOK, response.StatusCode())
}
