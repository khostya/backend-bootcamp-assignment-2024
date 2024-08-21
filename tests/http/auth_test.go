//go:build integration

package http

import (
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/api"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/api/models"
	"github.com/khostya/backend-bootcamp-assignment-2024/tests/http/client"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type (
	AuthTestSuite struct {
		suite.Suite
		ctx    context.Context
		client *api.Client
	}

	registerResult struct {
		Email    string
		ID       uuid.UUID
		Password string
	}

	userResult struct {
		Email    api.Email    `json:"email"`
		UserId   api.UserId   `json:"user_id"`
		UserType api.UserType `json:"user_type"`
	}
)

func TestAuth(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}

func (s *AuthTestSuite) SetupSuite() {
	s.T().Parallel()
	s.client = client.New()
	s.ctx = context.Background()
}

func (s *AuthTestSuite) SetupTest() {
	s.T().Parallel()
}

func (s *AuthTestSuite) TestDummyModerator() {
	s.dummy(model.Moderator)
}

func (s *AuthTestSuite) TestDummyClient() {
	s.dummy(model.Client)
}

func (s *AuthTestSuite) dummy(userType model.UserType) {
	resp, err := s.client.GetDummyLogin(s.ctx, &api.GetDummyLoginParams{UserType: userType})
	require.NoError(s.T(), err)

	parsedResp, err := api.ParseGetDummyLoginResponse(resp)
	require.NoError(s.T(), err)
	require.Equal(s.T(), http.StatusOK, parsedResp.StatusCode(), string(parsedResp.Body))
}

func getUser(ctx context.Context, t *testing.T, token string) userResult {
	client := client.New()

	resp, err := client.GetUser(ctx, addAuth(token))

	require.NoError(t, err)

	response, err := api.ParseGetUserResponse(resp)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode(), string(response.Body))

	return *response.JSON200
}

func addAuth(token string) func(ctx context.Context, req *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", "Bearer "+token)
		return nil
	}
}

func (s *AuthTestSuite) TestRegisterModerator() {
	register(s.ctx, s.T(), model.Moderator)
}

func (s *AuthTestSuite) TestRegisterClient() {
	register(s.ctx, s.T(), model.Client)
}

func (s *AuthTestSuite) TestLoginModerator() {
	login(s.ctx, s.T(), model.Moderator)
}

func (s *AuthTestSuite) TestLoginUser() {
	login(s.ctx, s.T(), model.Client)
}

func login(ctx context.Context, t *testing.T, userType model.UserType) (userResult, api.Token) {
	client := client.New()

	registerResult := register(ctx, t, userType)

	resp, err := client.PostLogin(ctx, api.PostLoginJSONRequestBody{
		Password: registerResult.Password,
		Id:       registerResult.ID,
	})
	require.NoError(t, err)

	response, err := api.ParsePostLoginResponse(resp)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode(), string(response.Body))

	user := getUser(ctx, t, *response.JSON200.Token)
	expected := userResult{UserId: registerResult.ID, Email: api.Email(registerResult.Email), UserType: userType}
	require.Equal(t, expected, user, string(response.Body))

	return expected, *response.JSON200.Token
}

func register(ctx context.Context, t *testing.T, userType model.UserType) registerResult {
	client := client.New()

	email := model.Email(gofakeit.Email())
	password := "414141432dssad"

	resp, err := client.PostRegister(ctx, api.PostRegisterJSONRequestBody{
		Password: password,
		UserType: userType,
		Email:    &email,
	})
	require.NoError(t, err)

	response, err := api.ParsePostRegisterResponse(resp)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode(), string(response.Body))

	return registerResult{
		Email:    string(email),
		ID:       *response.JSON200.UserId,
		Password: password,
	}
}
