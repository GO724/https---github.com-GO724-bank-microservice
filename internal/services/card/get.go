package card

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
)

func (b *cardService) Get(ctx context.Context, id uint) (*entity.Card, error) {
	fmt.Println("service: get card.id[%d]", id)
	return &entity.Card{}, nil
}
