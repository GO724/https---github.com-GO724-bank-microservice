package services

import (
	"bank-microservice/internal/repositories"
	srvBank "bank-microservice/internal/services/bank"
	srvCard "bank-microservice/internal/services/card"
	srvPerson "bank-microservice/internal/services/person"
)

type Services struct { // collection of services
	Card   srvCard.CardRepo
	Bank   srvBank.BankRepo
	Person srvPerson.PersonRepo
}

func (s *Services) New(repo *repositories.Repositories) { // init services by repository
	s.Card = srvCard.NewService(repo.Card)
	s.Bank = srvBank.NewService(repo.Bank)
	s.Person = srvPerson.NewService(repo.Person)
}
