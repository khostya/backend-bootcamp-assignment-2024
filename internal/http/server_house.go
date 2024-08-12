package http

import (
	"context"
	"github.com/go-chi/render"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/dto"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/api"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/middleware"
	"net/http"
)

type (
	houseUseCase interface {
		Create(ctx context.Context, param dto.CreateHouseParam) (domain.House, error)
		GetByID(ctx context.Context, id uint) (domain.House, error)
		Subscribe(ctx context.Context, id int, email string) error
	}
)

func (s *server) PostHouseCreate(w http.ResponseWriter, r *http.Request) {
	var req api.PostHouseCreateJSONBody
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status, house, err := s.postHouseCreate(r.Context(), req)
	if err != nil {
		s.error(w, r, status, err)
		return
	}

	s.json(w, r, http.StatusOK, house)
}

func (s *server) postHouseCreate(ctx context.Context, req api.PostHouseCreateJSONBody) (int, domain.House, error) {
	if ctx.Value(middleware.UserType) == domain.UserModerator {
		return http.StatusUnauthorized, domain.House{}, errUnauthorized
	}

	developer := ""
	if req.Developer != nil {
		developer = *req.Developer
	}

	param := dto.CreateHouseParam{Year: uint(req.Year), Address: req.Address, Developer: developer}
	err := s.validator.Struct(param)
	if err != nil {
		return http.StatusBadRequest, domain.House{}, err
	}

	house, err := s.useCases.House.Create(ctx, param)
	if err != nil {
		return http.StatusInternalServerError, domain.House{}, err
	}

	return http.StatusOK, house, nil
}

func (s *server) GetHouseId(w http.ResponseWriter, r *http.Request, id api.HouseId) {
	type flats struct {
		Flats []domain.Flat `json:"flats"`
	}

	house, ok := s.houseCache.Get(uint(id))
	if ok {
		s.json(w, r, http.StatusOK, flats{house.Flats})
		return
	}

	house, err := s.useCases.House.GetByID(r.Context(), uint(id))
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.houseCache.Put(uint(id), house)
	s.json(w, r, http.StatusOK, flats{house.Flats})
}

func (s *server) PostHouseIdSubscribe(w http.ResponseWriter, r *http.Request, id api.HouseId) {
	var req api.PostHouseIdSubscribeJSONBody
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.useCases.House.Subscribe(r.Context(), id, string(req.Email))
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
