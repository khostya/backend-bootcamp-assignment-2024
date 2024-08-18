package app

import (
	"context"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/cache"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/config"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/http"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/usecase"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/postgres"
)

const (
	databaseURL = "DATABASE_URL"
)

func Run(ctx context.Context, cfg config.Config) error {
	pool, err := postgres.NewPoolFromEnv(ctx, databaseURL)
	if err != nil {
		return err
	}
	defer pool.Close()

	deps := newDependencies(pool, cfg)

	useCases := usecase.NewUseCases(deps)

	cacheHouse := cfg.HTTP.Cache
	return <-http.MustRun(
		ctx,
		cfg.HTTP,
		http.Cache{
			House: cache.NewHouseCache(cacheHouse.House.Capacity, cacheHouse.House.TTL),
		},
		http.UseCases{
			Flat:  useCases.Flat,
			House: useCases.House,
			Auth:  useCases.Auth,
			User:  useCases.User,
		},
		deps.TokenManager,
	)
}
