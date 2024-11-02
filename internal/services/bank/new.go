package bank

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
)

func (b *bankService) New(ctx context.Context, bank entity.Bank) (*entity.Bank, error) {
	fmt.Println("service: new bank[%v]", bank)
	return &entity.Bank{}, nil
}
