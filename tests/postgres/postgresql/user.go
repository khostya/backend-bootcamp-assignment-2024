//go:build integration

package postgresql

import (
	"context"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo/schema"
)

func (d *DBPool) CreateUser(ctx context.Context, user domain.User) error {
	record := schema.NewUser(user)

	sql := `insert into bootcamp.users (password, email, type) VALUES ($1, $2, $3)`

	_, err := d.pool.Exec(ctx, sql, record.ID, record.Password, record.Email, record.Type)

	return err
}
