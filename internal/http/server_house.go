package http

import (
	"context"
	"github.com/go-chi/render"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/cache"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/dto"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/api"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/middleware"
	"net/http"
)

type (
	houseUseCase interface {
		Create(ctx context.Context, param dto.CreateHouseParam) (domain.House, error)
		GetByID(ctx context.Context, id uint, userType domain.UserType) (domain.House, error)
		Subscribe(ctx context.Context, subscription domain.Subscription) error
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
	if ctx.Value(middleware.KeyUserType) == domain.UserModerator {
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

	userType, ok := r.Context().Value(middleware.KeyUserType).(domain.UserType)
	if !ok || userType != domain.UserModerator {
		s.error(w, r, http.StatusUnauthorized, errUnauthorized)
		return
	}

	house, ok := s.houseCache.Get(cache.NewHouseKey(uint(id), userType))
	if ok {
		s.json(w, r, http.StatusOK, flats{house.Flats})
		return
	}

	house, err := s.useCases.House.GetByID(r.Context(), uint(id), userType)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	s.houseCache.Put(cache.NewHouseKey(uint(id), userType), house)
	s.json(w, r, http.StatusOK, flats{house.Flats})
}

func (s *server) PostHouseIdSubscribe(w http.ResponseWriter, r *http.Request, id api.HouseId) {
	isDummy, ok := r.Context().Value(middleware.KeyUserID).(bool)
	if !ok || isDummy {
		s.error(w, r, http.StatusUnauthorized, errUnauthorized)
		return
	}

	var req api.PostHouseIdSubscribeJSONBody
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.useCases.House.Subscribe(r.Context(), domain.Subscription{
		HouseID:   uint(id),
		UserEmail: string(req.Email),
	})
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
