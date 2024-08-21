// Package model provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package model

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for Status.
const (
	Approved     Status = "approved"
	Created      Status = "created"
	Declined     Status = "declined"
	OnModeration Status = "on moderation"
)

// Defines values for UserType.
const (
	Client    UserType = "client"
	Moderator UserType = "moderator"
)

// Address Адрес дома
type Address = string

// Date Дата + время
type Date = time.Time

// Developer Застройщик
type Developer = string

// Email Email пользователя
type Email = openapi_types.Email

// Flat Квартира
type Flat struct {
	// HouseId Идентификатор дома
	HouseId HouseId `json:"house_id"`

	// Id Идентификатор квартиры
	Id FlatId `json:"id"`

	// Price Цена квартиры в у.е.
	Price Price `json:"price"`

	// Rooms Количество комнат в квартире
	Rooms Rooms `json:"rooms"`

	// Status Статус квартиры
	Status Status `json:"status"`
}

// FlatId Идентификатор квартиры
type FlatId = int

// House Дом
type House struct {
	// Address Адрес дома
	Address Address `json:"address"`

	// CreatedAt Дата + время
	CreatedAt *Date `json:"created_at,omitempty"`

	// Developer Застройщик
	Developer *Developer `json:"developer"`

	// Id Идентификатор дома
	Id HouseId `json:"id"`

	// UpdateAt Дата + время
	UpdateAt *Date `json:"update_at,omitempty"`

	// Year Год постройки дома
	Year Year `json:"year"`
}

// HouseId Идентификатор дома
type HouseId = int

// Password Пароль пользователя
type Password = string

// Price Цена квартиры в у.е.
type Price = int

// Rooms Количество комнат в квартире
type Rooms = int

// Status Статус квартиры
type Status string

// Token Авторизационный токен
type Token = string

// UserId Идентификатор пользователя
type UserId = openapi_types.UUID

// UserType Тип пользователя
type UserType string

// Year Год постройки дома
type Year = int

// N5xx defines model for 5xx.
type N5xx struct {
	// Code Код ошибки. Предназначен для классификации проблем и более быстрого решения проблем.
	Code *int `json:"code,omitempty"`

	// Message Описание ошибки
	Message string `json:"message"`

	// RequestId Идентификатор запроса. Предназначен для более быстрого поиска проблем.
	RequestId *string `json:"request_id,omitempty"`
}
