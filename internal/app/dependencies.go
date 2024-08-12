package app

import (
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/config"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo/transactor"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/usecase"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/auth"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/hash"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/postgres"
)

func newDependencies(pool *postgres.Pool, cfg config.Config) usecase.Dependencies {
	transactor := transactor.NewTransactionManager(pool)
	pgRepositories := repo.NewRepositories(transactor)
	return usecase.Dependencies{
		Pg:         pgRepositories,
		Transactor: transactor,

		TokenManager:   auth.NewManager(cfg.Auth.SigningKey),
		AccessTokenTTL: cfg.Auth.AccessTokenTTL,
		PasswordHasher: hash.NewPasswordHasher(cfg.Auth.PasswordCostBcrypt),
	}
}
