package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var emailRegexp = regexp.MustCompile("\\S+@\\S+\\.\\S+")

func (v *Validator) emailValidate(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}

	return emailRegexp.MatchString(fl.Field().String())
}
