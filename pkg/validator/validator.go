package validator

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	v         *validator.Validate
	passwdErr error
}

func NewValidate() (*Validator, error) {
	v := Validator{v: validator.New(validator.WithRequiredStructEnabled())}

	err := v.v.RegisterValidation("email", v.emailValidate)

	return &v, err
}

func (v *Validator) Struct(i any) error {
	return v.v.Struct(i)
}
