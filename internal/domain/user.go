package domain

import (
	"backend-bootcamp-assignment-2024/internal/dto"
	"github.com/google/uuid"
)

type UserType string

const (
	UserModerator UserType = "moderator"
	UserClient    UserType = "client"
)

type (
	Token string

	User struct {
		ID       uuid.UUID
		Email    string
		UserType UserType
		Password string
	}
)

func NewUser(param dto.RegisterUserParam, hashedPassword string) User {
	return User{
		ID:       uuid.New(),
		Email:    param.Email,
		UserType: UserType(param.UserType),
		Password: hashedPassword,
	}
}
