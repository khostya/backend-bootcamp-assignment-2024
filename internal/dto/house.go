package dto

type CreateHouseParam struct {
	Address   string `validate:"required,min=1"`
	Year      uint   `validate:"required,min=1"`
	Developer string
}
