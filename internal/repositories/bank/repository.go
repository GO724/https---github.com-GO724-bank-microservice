package bank

import (
	"bank-microservice/internal/entity"
	"bank-microservice/internal/infrastructure/database"
	"context"
)

type Bank interface {
	New(ctx context.Context, b entity.Bank) error                   // New record
	Get(ctx context.Context, bic uint) (entity.Bank, error)         // Read record
	Set(ctx context.Context, b entity.Bank) error                   // Update record
	List(ctx context.Context, b entity.Bank) ([]entity.Bank, error) // List of record filtered on b(Bank)
	Delete(ctx context.Context, bic uint) error                     // Delete record
	Migrate(ctx context.Context) error                              // Create table & fill demo record
}

type bankRepository struct {
	db *database.Database
}

func NewRepository(db *database.Database) *bankRepository {
	return &bankRepository{
		db: db,
	}
}

//TODO: пофиксить все функции, и вынести их в отдельные файлы одноименные
