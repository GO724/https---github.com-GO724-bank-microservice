// Сущности слоя БД

package repositories

import (
	"bank-microservice/internal/infrastructure/database"
	repoBank "bank-microservice/internal/repositories/bank"
	repoCard "bank-microservice/internal/repositories/card"
	repoPerson "bank-microservice/internal/repositories/person"
)

type Repositories struct { // collection of repositories
	Card   repoCard.Card
	Bank   repoBank.Bank
	Person repoPerson.Person
}

func NewRepository(db *database.Database) *Repositories {
	return &Repositories{
		Card:   repoCard.NewRepository(db),
		Bank:   repoBank.NewRepository(db),
		Person: repoPerson.NewRepository(db),
	}
}
