package bank

import (
	"bank-microservice/internal/entity"
	"context"
)

func (b bankRepository) List(ctx context.Context, bank entity.Bank) ([]entity.Bank, error) {
	//	// List of record filtered on b(Bank)
	//	sb := make([]entity.Bank, 0, 10)
	//	return sb, nil
	return nil, nil
}
