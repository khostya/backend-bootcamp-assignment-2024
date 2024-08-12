package repo

import (
	"backend-bootcamp-assignment-2024/internal/domain"
	"backend-bootcamp-assignment-2024/internal/repo/exec"
	"backend-bootcamp-assignment-2024/internal/repo/schema"
	"backend-bootcamp-assignment-2024/internal/repo/transactor"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const (
	userTable = "bootcamp.users"
)

type (
	User struct {
		queryEngineProvider transactor.QueryEngineProvider
	}
)

func (u User) GetByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	db := u.queryEngineProvider.GetQueryEngine(ctx)

	query := sq.Select(schema.User{}.Columns()...).
		From(userTable).
		Where("id = $1", id).
		PlaceholderFormat(sq.Dollar)

	user, err := exec.ScanOne[schema.User](ctx, query, db)
	if err != nil {
		return domain.User{}, err
	}

	return schema.NewDomainUser(user), nil
}

func (u User) Create(ctx context.Context, user domain.User) error {
	db := u.queryEngineProvider.GetQueryEngine(ctx)

	record := schema.NewUser(user)
	query := sq.Insert(userTable).
		Columns(record.Columns()...).
		Values(record.Values()...).
		PlaceholderFormat(sq.Dollar)

	return exec.Insert(ctx, query, db)
}

func NewUserRepo(provider transactor.QueryEngineProvider) User {
	return User{provider}
}
