package bank

import (
	"bank-microservice/internal/entity"
	repo "bank-microservice/internal/repositories/bank"
	"context"
)

type BankRepo interface {
	New(ctx context.Context, c entity.Bank) (*entity.Bank, error)    // New record
	Get(ctx context.Context, id uint) (*entity.Bank, error)          // Read record
	Set(ctx context.Context, c entity.Bank) (*entity.Bank, error)    // Update record
	Delete(ctx context.Context, id uint) error                       // Delete record
	List(ctx context.Context, c entity.Bank) ([]*entity.Bank, error) // List of record filtered on c(Bank)
}

// type Bank interface {
// }

type bankService struct {
	bankRepo repo.Bank
}

func NewService(rb repo.Bank) *bankService {
	return &bankService{
		bankRepo: rb,
	}
}
