package app

import (
	"backend-bootcamp-assignment-2024/internal/config"
	"backend-bootcamp-assignment-2024/internal/http"
	"backend-bootcamp-assignment-2024/internal/usecase"
	"backend-bootcamp-assignment-2024/pkg/postgres"
	"context"
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

	return <-http.MustRun(ctx, cfg.HTTP, http.UseCases{
		Flat:  useCases.Flat,
		House: useCases.House,
		Auth:  useCases.Auth,
	}, deps.TokenManager)
}
