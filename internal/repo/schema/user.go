package schema

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
)

type (
	User struct {
		ID       uuid.UUID        `db:"id"`
		Email    sql.Null[string] `db:"email"`
		Type     string           `db:"type"`
		Password string           `db:"password"`
	}
)

func (u User) Values() []any {
	return []any{u.ID, u.Email, u.Password, u.Type}
}

func (u User) Columns() []string {
	return []string{"id", "email", "password", "type"}
}

func NewUser(user domain.User) User {
	return User{
		ID:       user.ID,
		Password: user.Password,
		Type:     string(user.UserType),
		Email:    sql.Null[string]{V: user.Email, Valid: user.Email != ""},
	}
}

func NewDomainUser(user User) domain.User {
	return domain.User{
		ID:       user.ID,
		Password: user.Password,
		UserType: domain.UserType(user.Type),
		Email:    user.Email.V,
	}
}
