package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher struct {
	cost int
}

func NewPasswordHasher(cost int) PasswordHasher {
	return PasswordHasher{cost: cost}
}

func (h PasswordHasher) Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}
	return string(hashed), err
}

func (h PasswordHasher) Equals(hashed string, password string) bool {
	eq := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return eq == nil
}
