package domain

import (
	"backend-bootcamp-assignment-2024/internal/dto"
	"github.com/google/uuid"
)

const (
	FlatCreated    FlatStatus = "created"
	FlatApproved   FlatStatus = "approved"
	FlatDeclined   FlatStatus = "declined"
	FlatModeration FlatStatus = "on moderation"
)

type (
	FlatStatus string

	Flat struct {
		ID          uint       `json:"id"`
		HouseID     uint       `json:"house_id"`
		Price       uint       `json:"price"`
		Rooms       uint       `json:"rooms"`
		Status      FlatStatus `json:"status"`
		ModeratorID uuid.UUID  `json:"_"`
	}
)

func NewFlat(param dto.CreateFlatParam) Flat {
	return Flat{
		HouseID: param.HouseID,
		Price:   param.Price,
		Rooms:   param.Rooms,
		Status:  FlatCreated,
	}
}
