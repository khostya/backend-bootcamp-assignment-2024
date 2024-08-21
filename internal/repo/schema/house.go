package schema

import (
	"database/sql"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"time"
)

type (
	House struct {
		ID uint `db:"h_house_id"`

		Address   string           `db:"address"`
		Year      uint             `db:"year"`
		Developer sql.Null[string] `db:"developer"`

		CreatedAt       time.Time `db:"created_at"`
		LastFlatAddedAt time.Time `db:"last_flat_added_at"`
	}

	FlatHouse struct {
		Flat
		House
	}
)

func (h House) SelectColumns() []string {
	return []string{"houses.id as h_house_id", "address", "year", "developer", "created_at", "last_flat_added_at"}
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

func NewDomainHouse(house House) domain.House {
	return domain.House{
		ID:              house.ID,
		Address:         house.Address,
		Developer:       house.Developer.V,
		Year:            house.Year,
		CreatedAt:       house.CreatedAt,
		LastFlatAddedAt: house.LastFlatAddedAt,
	}
}

func NewDomainHouseWithFlats(flatsHouse []FlatHouse) domain.House {
	house := NewDomainHouse(flatsHouse[0].House)

	for _, flat := range flatsHouse {
		house.Flats = append(house.Flats, NewDomainFlat(flat.Flat))
	}

	return house
}
