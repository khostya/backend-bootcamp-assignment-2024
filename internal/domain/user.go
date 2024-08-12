package domain

import (
	"github.com/google/uuid"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/dto"
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
