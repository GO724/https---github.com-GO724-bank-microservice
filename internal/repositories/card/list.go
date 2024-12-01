package card

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
)

func (c cardRepository) List(ctx context.Context, card entity.Card) ([]entity.Card, error) {
	//	// List of record filtered on b(Bank)
	//	sb := make([]entity.Bank, 0, 10)
	//	return sb, nil
	fmt.Println("service: list all of card[%v]\n", card)
	return nil, nil
}
