package card

import (
	"bank-microservice/internal/entity"
	"bank-microservice/internal/infrastructure/database"
	"context"
)

type Card interface {
	New(ctx context.Context, c entity.Card) error                   // New record
	Get(ctx context.Context, id uint) (entity.Card, error)          // Read record
	Set(ctx context.Context, c entity.Card) error                   // Update record
	List(ctx context.Context, c entity.Card) ([]entity.Card, error) // List of record filtered on c(Card)
	Delete(ctx context.Context, id uint) error                      // Delete record
	Migrate(ctx context.Context) error                              // Create table & fill demo record
}

type cardRepository struct {
	db *database.Database
}

func NewRepository(db *database.Database) *cardRepository {
	return &cardRepository{
		db: db,
	}
}

//TODO: пофиксить все функции, и вынести их в отдельные файлы одноименные
