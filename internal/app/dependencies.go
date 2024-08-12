package app

import (
	"backend-bootcamp-assignment-2024/internal/config"
	"backend-bootcamp-assignment-2024/internal/repo"
	"backend-bootcamp-assignment-2024/internal/repo/transactor"
	"backend-bootcamp-assignment-2024/internal/usecase"
	"backend-bootcamp-assignment-2024/pkg/auth"
	"backend-bootcamp-assignment-2024/pkg/hash"
	"backend-bootcamp-assignment-2024/pkg/postgres"
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
