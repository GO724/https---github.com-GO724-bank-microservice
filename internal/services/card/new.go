package card

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
)

func (b *cardService) New(ctx context.Context, card entity.Card) (*entity.Card, error) {
	fmt.Println("service: new card[%v]", card)
	return &entity.Card{}, nil
}
