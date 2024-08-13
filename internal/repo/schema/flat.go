package schema

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
)

type (
	Flat struct {
		ID          uint   `db:"flat_id"`
		HouseID     uint   `db:"house_id"`
		Price       uint   `db:"price"`
		Rooms       uint   `db:"rooms"`
		Status      string `db:"flat_status"`
		ModeratorID sql.Null[uuid.UUID]
	}
)

func (f Flat) SelectColumns() []string {
	return []string{"flats.id as flat_id", "house_id", "price", "rooms", "flats.status as flat_status", "moderator_id"}
}

func (f Flat) InsertValues() []any {
	return []any{f.HouseID, f.Price, f.Rooms, f.Status, f.ModeratorID}
}

func (f Flat) InsertColumns() []string {
	return []string{"house_id", "price", "rooms", "status", "moderator_id"}
}

func NewFlat(flat domain.Flat) Flat {
	return Flat{
		ID:          flat.ID,
		HouseID:     flat.HouseID,
		Price:       flat.Price,
		Rooms:       flat.Rooms,
		Status:      string(flat.Status),
		ModeratorID: sql.Null[uuid.UUID]{V: flat.ModeratorID, Valid: flat.ModeratorID.String() != uuid.UUID{}.String()},
	}
}

func NewDomainFlat(flat Flat) domain.Flat {
	return domain.Flat{
		ID:          flat.ID,
		HouseID:     flat.HouseID,
		Price:       flat.Price,
		Rooms:       flat.Rooms,
		Status:      domain.FlatStatus(flat.Status),
		ModeratorID: flat.ModeratorID.V,
	}
}
