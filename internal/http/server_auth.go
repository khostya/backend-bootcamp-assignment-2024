package http

import (
	"backend-bootcamp-assignment-2024/internal/domain"
	"backend-bootcamp-assignment-2024/internal/dto"
	"backend-bootcamp-assignment-2024/internal/http/api"
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
)

type (
	authUseCase interface {
		Login(ctx context.Context, param dto.LoginUserParam) (domain.Token, error)
		Register(ctx context.Context, param dto.RegisterUserParam) (uuid.UUID, error)
		DummyLogin(ctx context.Context, userType domain.UserType) (domain.Token, error)
	}
)

func (s *server) PostLogin(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token domain.Token `json:"token"`
	}

	var req api.PostLoginJSONBody
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	status, token, err := s.postLogin(ctx, req)
	if err != nil {
		s.error(w, r, status, err)
		return
	}

	s.json(w, r, status, response{Token: token})
}

func (s *server) postLogin(ctx context.Context, req api.PostLoginJSONBody) (int, domain.Token, error) {
	param := dto.LoginUserParam{Id: req.Id, Password: req.Password}
	err := s.validator.Struct(param)
	if err != nil {
		return http.StatusBadRequest, "", err
	}

	token, err := s.useCases.Auth.Login(ctx, param)
	if err != nil {
		return http.StatusInternalServerError, token, err
	}

	return http.StatusOK, token, nil
}

func (s *server) PostRegister(w http.ResponseWriter, r *http.Request) {
	type response struct {
		UserId uuid.UUID `json:"user_id"`
	}

	var req api.PostRegisterJSONBody
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status, id, err := s.postRegister(r.Context(), req)
	if err != nil {
		s.error(w, r, status, err)
		return
	}

	s.json(w, r, status, response{id})
}

func (s *server) postRegister(ctx context.Context, req api.PostRegisterJSONBody) (int, uuid.UUID, error) {
	email := ""
	if req.Email != nil {
		email = string(*req.Email)
	}

	param := dto.RegisterUserParam{Email: email, Password: req.Password, UserType: string(req.UserType)}
	err := s.validator.Struct(param)
	if err != nil {
		return http.StatusBadRequest, uuid.UUID{}, err
	}

	id, err := s.useCases.Auth.Register(ctx, param)
	if err != nil {
		return http.StatusInternalServerError, id, err
	}

	return http.StatusOK, id, nil
}

func (s *server) GetDummyLogin(w http.ResponseWriter, r *http.Request, params api.GetDummyLoginParams) {
	type response struct {
		Token api.Token `json:"token"`
	}

	ctx := r.Context()
	token, err := s.useCases.Auth.DummyLogin(ctx, domain.UserType(params.UserType))
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	resp := response{Token: api.Token(token)}
	s.json(w, r, http.StatusOK, resp)
}
