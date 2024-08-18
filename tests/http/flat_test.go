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
	"testing"
)

type (
	FlatTestSuite struct {
		suite.Suite
		ctx    context.Context
		client *api.Client
	}
)

func TestFlat(t *testing.T) {
	suite.Run(t, new(FlatTestSuite))
}

func (s *FlatTestSuite) SetupSuite() {
	s.T().Parallel()
	s.client = client.New()
	s.ctx = context.Background()
}

func (s *FlatTestSuite) SetupTest() {
	s.T().Parallel()
}

func (s *FlatTestSuite) TestCreate() {
	createFlat(s.ctx, s.T())
}

func createFlat(ctx context.Context, t *testing.T) (*api.House, api.Flat, api.Token) {
	client := client.New()

	house, _, token := createHouse(ctx, t)

	flatReq := api.PostFlatCreateJSONRequestBody{
		HouseId: house.Id,
		Price:   int(gofakeit.Uint8()),
		Rooms:   int(gofakeit.Uint8()),
	}

	resp, err := client.PostFlatCreate(ctx, flatReq, addAuth(token))
	require.NoError(t, err)

	response, err := api.ParsePostFlatCreateResponse(resp)
	require.NoError(t, err)

	flat := api.Flat{
		Id:      response.JSON200.Id,
		HouseId: flatReq.HouseId,
		Price:   flatReq.Price,
		Rooms:   flatReq.Rooms,
		Status:  model.Created,
	}
	require.Equal(t, flat, *response.JSON200)

	return house, flat, token
}

func (s *FlatTestSuite) TestGet() {
	house, flat, token := createFlat(s.ctx, s.T())

	resp, err := s.client.GetHouseId(s.ctx, house.Id, addAuth(token))
	require.NoError(s.T(), err)

	response, err := api.ParseGetHouseIdResponse(resp)
	require.NoError(s.T(), err)

	require.Equal(s.T(), []api.Flat{flat}, response.JSON200.Flats)
}

func (s *FlatTestSuite) TestUpdateStatus() {
	_, flat, token := createFlat(s.ctx, s.T())

	resp, err := s.client.PostFlatUpdate(s.ctx, api.PostFlatUpdateJSONRequestBody{
		Id:     flat.Id,
		Status: model.OnModeration,
	}, addAuth(token))
	require.NoError(s.T(), err)

	response, err := api.ParsePostFlatUpdateResponse(resp)
	require.NoError(s.T(), err)

	require.Equal(s.T(), model.OnModeration, response.JSON200.Status)
}
