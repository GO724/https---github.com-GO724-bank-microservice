package card

import (
	"bank-microservice/internal/entity"
	"context"
)

func (c cardRepository) List(ctx context.Context, card entity.Card) ([]entity.Card, error) {
	//	// List of record filtered on b(Bank)
	//	sb := make([]entity.Bank, 0, 10)
	//	return sb, nil
	return nil, nil
}
