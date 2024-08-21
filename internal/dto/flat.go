package dto

import "github.com/google/uuid"

type (
	CreateFlatParam struct {
		HouseID uint `validate:"required,min=1"`
		Price   uint `validate:"required,min=1"`
		Rooms   uint `validate:"required,min=1"`
	}

	FlatCreateResult struct {
		ID, Number uint
	}

	UpdateFlatParam struct {
		Id          uint      `validate:"required"`
		Status      string    `validate:"required"`
		ModeratorID uuid.UUID `validate:"required"`
	}
)
