package person

import (
	"bank-microservice/internal/entity"
	"bank-microservice/internal/infrastructure/database"
	"context"
)

type Person interface {
	New(ctx context.Context, p entity.Person) error                     // New record
	Get(ctx context.Context, inn uint) (entity.Person, error)           // Read record
	Set(ctx context.Context, p entity.Person) error                     // Update record
	List(ctx context.Context, p entity.Person) ([]entity.Person, error) // List of record filtered on b(Bank)
	Delete(ctx context.Context, inn uint) error                         // Delete record
	Migrate(ctx context.Context) error                                  // Create table & fill demo record
}

type personRepository struct {
	db *database.Database
}

func NewRepository(db *database.Database) *personRepository {
	return &personRepository{
		db: db,
	}
}

//READY:TODO: пофиксить все функции, и вынести их в отдельные файлы одноименные
