package card

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
)

func (b *cardService) Set(ctx context.Context, card entity.Card) (*entity.Card, error) {
	fmt.Println("service: set card[%v]", card)
	return &entity.Card{}, nil
}
