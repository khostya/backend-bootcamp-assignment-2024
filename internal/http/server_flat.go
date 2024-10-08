package http

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/dto"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/api"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http/middleware"
	"net/http"
)

type (
	flatUseCase interface {
		Create(ctx context.Context, param dto.CreateFlatParam) (domain.Flat, error)
		Update(ctx context.Context, param dto.UpdateFlatParam) (domain.Flat, error)
	}
)

func (s *server) PostFlatCreate(w http.ResponseWriter, r *http.Request) {
	var req api.PostFlatCreateJSONBody
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status, flat, err := s.postFlatCreate(r.Context(), req)
	if err != nil {
		s.error(w, r, status, err)
		return
	}

	s.json(w, r, http.StatusOK, flat)
}

func (s *server) postFlatCreate(ctx context.Context, req api.PostFlatCreateJSONBody) (int, domain.Flat, error) {
	param := dto.CreateFlatParam{Rooms: uint(req.Rooms), Price: uint(req.Price), HouseID: uint(req.HouseId)}
	err := s.validator.Struct(param)
	if err != nil {
		return http.StatusBadRequest, domain.Flat{}, err
	}

	flat, err := s.useCases.Flat.Create(ctx, param)
	if err != nil {
		return http.StatusInternalServerError, flat, err
	}

	return http.StatusOK, flat, nil
}

func (s *server) PostFlatUpdate(w http.ResponseWriter, r *http.Request) {
	var req api.PostFlatUpdateJSONBody
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status, flat, err := s.postFlatUpdate(r.Context(), req)
	if err != nil {
		s.error(w, r, status, err)
		return
	}

	s.json(w, r, http.StatusOK, flat)
}

func (s *server) postFlatUpdate(ctx context.Context, req api.PostFlatUpdateJSONBody) (int, domain.Flat, error) {
	userType, ok := ctx.Value(middleware.KeyUserType).(domain.UserType)
	if !ok || userType != domain.UserModerator {
		return http.StatusUnauthorized, domain.Flat{}, errUnauthorized
	}

	moderatorID, ok := ctx.Value(middleware.KeyUserID).(uuid.UUID)
	if !ok {
		return http.StatusUnauthorized, domain.Flat{}, errUnauthorized
	}

	param := dto.UpdateFlatParam{Id: uint(req.Id), Status: string(req.Status), ModeratorID: moderatorID}
	err := s.validator.Struct(param)
	if err != nil {
		return http.StatusBadRequest, domain.Flat{}, err
	}

	flat, err := s.useCases.Flat.Update(ctx, param)
	if err != nil {
		return http.StatusInternalServerError, domain.Flat{}, err
	}

	return http.StatusOK, flat, nil
}
