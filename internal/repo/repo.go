package repo

import (
	"backend-bootcamp-assignment-2024/internal/repo/transactor"
)

type Repositories struct {
	Flat  Flat
	House House
	User  User
}

func NewRepositories(provider transactor.QueryEngineProvider) Repositories {
	return Repositories{
		Flat:  NewFlatRepo(provider),
		User:  NewUserRepo(provider),
		House: NewHouseRepo(provider),
	}
}
