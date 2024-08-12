//go:build integration

package postgres

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"time"
)

func NewUser() domain.User {
	return domain.User{
		Email:    gofakeit.Email(),
		ID:       uuid.New(),
		UserType: domain.UserClient,
		Password: gofakeit.Password(true, true, true, true, true, 10),
	}
}

func NewHouses() domain.House {
	return domain.House{
		ID:              0,
		Address:         "3131",
		Year:            3131,
		Developer:       "3131",
		CreatedAt:       time.Now(),
		LastFlatAddedAt: time.Now(),
	}
}

func NewFlats(houseID uint) domain.Flat {
	return domain.Flat{
		ID:      0,
		HouseID: houseID,
		Price:   1,
		Rooms:   4,
		Status:  domain.FlatCreated,
	}
}
