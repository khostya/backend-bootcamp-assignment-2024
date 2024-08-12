package usecase

import (
	"context"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo/transactor"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/auth"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/hash"
	"time"
)

type (
	transactionManager interface {
		RunRepeatableRead(ctx context.Context, fx func(ctxTX context.Context) error) error
		Unwrap(err error) error
	}

	Dependencies struct {
		Pg         repo.Repositories
		Transactor *transactor.TransactionManager

		PasswordHasher hash.PasswordHasher
		TokenManager   auth.TokenManager
		AccessTokenTTL time.Duration
	}

	UseCases struct {
		Flat  Flat
		House House
		Deps  Dependencies
		Auth  Auth
	}
)

func NewUseCases(deps Dependencies) UseCases {
	pg := deps.Pg
	transactor := deps.Transactor

	return UseCases{
		Deps:  deps,
		Flat:  NewFlatUseCase(pg.Flat, pg.House, transactor),
		House: NewHouseUseCase(pg.House, transactor),
		Auth: NewAuthUseCase(AuthDeps{
			TransactionManager: deps.Transactor,
			UserRepo:           deps.Pg.User,
			PasswordHasher:     deps.PasswordHasher,
			TokenManager:       deps.TokenManager,
			AccessTokenTTL:     deps.AccessTokenTTL,
		}),
	}
}
