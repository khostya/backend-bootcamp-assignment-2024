package domain

import (
	"backend-bootcamp-assignment-2024/internal/dto"
	"time"
)

type House struct {
	ID uint `json:"id"`

	Address   string `json:"address"`
	Year      uint   `json:"year"`
	Developer string `json:"developer"`

	Flats []Flat `json:"_"`

	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func NewHouse(param dto.CreateHouseParam) House {
	return House{
		ID:        0,
		Year:      param.Year,
		Address:   param.Address,
		Developer: param.Developer,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}
