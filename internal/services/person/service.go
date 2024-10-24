package person

import "bank-microservice/internal/repositories/person"

type Person interface {
}

type personService struct {
	person person.Person // repository
}

func NewService(p person.Person) Person {
	return &personService{
		person: p,
	}
}
