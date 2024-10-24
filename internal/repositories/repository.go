// Сущности слоя БД

package repositories

import (
	"bank-microservice/internal/infrastructure/database"
	"bank-microservice/internal/repositories/bank"
	"bank-microservice/internal/repositories/card"
	"bank-microservice/internal/repositories/person"
)

type Repositories struct {
	Card   card.Card
	Bank   bank.Bank
	Person person.Person
}

func NewRepository(db *database.Database) Repositories {
	return Repositories{
		Card:   card.NewRepository(db),
		Bank:   bank.NewRepository(db),
		Person: person.NewRepository(db),
	}
}
