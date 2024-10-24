package services

import (
	"bank-microservice/internal/repositories"
	"bank-microservice/internal/services/bank"
	"bank-microservice/internal/services/card"
	"bank-microservice/internal/services/person"
)

type Services struct { // collection of services
	Card   card.Card
	Bank   bank.Bank
	Person person.Person
}

func (s *Services) New(repo *repositories.Repositories) error { // init services by repository
	s.Card = card.NewService(repo.Card)
	s.Bank = bank.NewService(repo.Bank)
	s.Person = person.NewService(repo.Person)
	return nil
}
