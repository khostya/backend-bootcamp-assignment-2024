package http

import (
	"backend-bootcamp-assignment-2024/internal/cache"
	"backend-bootcamp-assignment-2024/internal/http/api"
	"backend-bootcamp-assignment-2024/pkg/validator"
	"errors"
)

var _ api.ServerInterface = (*server)(nil)

var errUnauthorized = errors.New("unauthorized")

type (
	UseCases struct {
		Flat  flatUseCase
		House houseUseCase
		Auth  authUseCase
	}

	Cache struct {
		House *cache.HouseCache
	}

	server struct {
		useCases   UseCases
		validator  *validator.Validator
		houseCache *cache.HouseCache
	}
)

func newServer(useCases UseCases, cache Cache) (*server, error) {
	validator, err := validator.NewValidate()
	if err != nil {
		return nil, err
	}

	return &server{
		useCases:   useCases,
		validator:  validator,
		houseCache: cache.House,
	}, nil
}
