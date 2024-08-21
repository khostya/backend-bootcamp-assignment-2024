package domain

import (
	"github.com/google/uuid"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/dto"
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
		Number      uint       `json:"-"`
		ModeratorID uuid.UUID  `json:"-"`
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
