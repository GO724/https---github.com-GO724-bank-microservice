package bank

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
)

func (b *bankService) Set(ctx context.Context, bank entity.Bank) (*entity.Bank, error) {
	fmt.Println("service: set bank[%v]", bank)
	return &entity.Bank{}, nil
}
