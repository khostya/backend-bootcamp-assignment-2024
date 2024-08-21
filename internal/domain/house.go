package domain

import (
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/dto"
	"time"
)

type House struct {
	ID uint `json:"id"`

	Address   string `json:"address"`
	Year      uint   `json:"year"`
	Developer string `json:"developer"`

	Flats []Flat `json:"-"`

	CreatedAt       time.Time `json:"created_at"`
	LastFlatAddedAt time.Time `json:"update_at"`
}

func NewHouse(param dto.CreateHouseParam) House {
	return House{
		ID:              0,
		Year:            param.Year,
		Address:         param.Address,
		Developer:       param.Developer,
		CreatedAt:       time.Now(),
		LastFlatAddedAt: time.Now(),
	}
}
