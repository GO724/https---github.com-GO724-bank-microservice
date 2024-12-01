package person

import (
	"bank-microservice/internal/entity"
	repo "bank-microservice/internal/repositories/person"
	"context"
)

type PersonRepo interface {
	New(ctx context.Context, c entity.Person) (*entity.Person, error)    // New record
	Get(ctx context.Context, id uint) (*entity.Person, error)            // Read record
	Set(ctx context.Context, c entity.Person) (*entity.Person, error)    // Update record
	Delete(ctx context.Context, id uint) error                           // Delete record
	List(ctx context.Context, c entity.Person) ([]*entity.Person, error) // List of record filtered on c(Person)

}

type personService struct {
	person repo.Person // repository
}

func NewService(rp repo.Person) *personService {
	return &personService{
		person: rp,
	}
}
