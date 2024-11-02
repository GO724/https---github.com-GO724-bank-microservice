package bank

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
)

func (b *bankService) Get(ctx context.Context, bic uint) (*entity.Bank, error) {
	fmt.Println("service: get bank.bic[%d]", bic)
	return &entity.Bank{}, nil
}
