package schema

import (
	"database/sql"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"time"
)

type (
	House struct {
		ID uint `db:"id"`

		Address   string           `db:"address"`
		Year      uint             `db:"year"`
		Developer sql.Null[string] `db:"developer"`

		CreatedAt       time.Time `db:"created_at"`
		LastFlatAddedAt time.Time `db:"last_flat_added_at"`
	}
)

func (h House) Values() []any {
	return []any{h.ID, h.Address, h.Year, h.Developer, h.CreatedAt, h.LastFlatAddedAt}
}

func (h House) Columns() []string {
	return []string{"id", "address", "year", "developer", "created_at", "updated_at", "last_flat_added_at"}
}

func (h House) ValuesInsert() []any {
	return []any{h.Address, h.Year, h.Developer, h.CreatedAt, h.LastFlatAddedAt}
}

func (h House) ColumnsInsert() []string {
	return []string{"address", "year", "developer", "created_at", "last_flat_added_at"}
}

func NewHouse(user domain.House) House {
	return House{
		ID:              user.ID,
		Address:         user.Address,
		Developer:       sql.Null[string]{V: user.Developer, Valid: user.Developer != ""},
		Year:            user.Year,
		CreatedAt:       user.CreatedAt,
		LastFlatAddedAt: user.LastFlatAddedAt,
	}
}

func NewDomainHouse(user House) domain.House {
	return domain.House{
		ID:              user.ID,
		Address:         user.Address,
		Developer:       user.Developer.V,
		Year:            user.Year,
		CreatedAt:       user.CreatedAt,
		LastFlatAddedAt: user.LastFlatAddedAt,
	}
}
