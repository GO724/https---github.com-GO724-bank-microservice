package bank

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
)

func (b bankRepository) List(ctx context.Context, bank entity.Bank) ([]entity.Bank, error) {
	//	// List of record filtered on b(Bank)
	//	sb := make([]entity.Bank, 0, 10)
	//	return sb, nil
	fmt.Println("service: list all of bank[%v]\n", bank)
	return nil, nil
}
