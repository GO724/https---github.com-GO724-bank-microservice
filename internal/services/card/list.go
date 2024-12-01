package card

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
)

func (b *cardService) List(ctx context.Context, card entity.Card) ([]*entity.Card, error) {
	fmt.Println("service: list card[%v]", card)
	return make([]*entity.Card, 0), nil
}
