package usecase

import (
	"context"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo/transactor"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/auth"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/hash"
	emailsender "github.com/khostya/backend-bootcamp-assignment-2024/pkg/sender"
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

		Sender *emailsender.Sender
	}

	UseCases struct {
		Flat  Flat
		House House
		Deps  Dependencies
		Auth  Auth
		User  User
	}
)

func NewUseCases(deps Dependencies) UseCases {
	pg := deps.Pg
	transactor := deps.Transactor

	return UseCases{
		Deps:  deps,
		Flat:  NewFlatUseCase(pg.Flat, pg.House, pg.Subscription, transactor, deps.Sender),
		House: NewHouseUseCase(pg.House, pg.Subscription, transactor),
		Auth: NewAuthUseCase(AuthDeps{
			TransactionManager: deps.Transactor,
			UserRepo:           deps.Pg.User,
			PasswordHasher:     deps.PasswordHasher,
			TokenManager:       deps.TokenManager,
			AccessTokenTTL:     deps.AccessTokenTTL,
		}),
		User: NewUserUseCase(pg.User),
	}
}
