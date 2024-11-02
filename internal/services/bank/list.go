package bank

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
)

func (b *bankService) List(ctx context.Context, bank entity.Bank) ([]*entity.Bank, error) {
	fmt.Println("service: get bank[%v]", bank)
	return make([]*entity.Bank, 0), nil
}
