package dto

import "github.com/google/uuid"

type RegisterUserParam struct {
	Email    string `validate:"email"`
	Password string `validate:"required"`
	UserType string `validate:"required"`
}

type LoginUserParam struct {
	Id       uuid.UUID `validate:"required"`
	Password string    `validate:"required"`
}
