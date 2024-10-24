package bank

import (
	"context"
)

func (b *bankService) Delete(ctx context.Context, bic uint) error {
	// bank, err := b.bankRepo.Get(ctx, bic)
	// if err != nil {
	// 	return fmt.Errorf("can't delete a bank.bic[%d]: read error: %v", bic, err)
	// }

	// some logic: can't delete a bank if cards from this bank exist

	// походу не лучшая идея передавать параметры поиска в структуре сущности - нельзя nil на игнорируемых в поиске полях. мапа или массив/слайс?

	// searchCard := entity.Card{Id: 0,
	// 	Person:  entity.Person{Inn: 0, Name: ""},
	// 	Bank:    bank,
	// 	Expires: time.Date(10000, time.January, 1, 0, 0, 0, 0, &time.Location{}),
	// }

	// // как в логике использовать другие сущности?

	// listCard, err := card.Card.List(ctx, searchCard)
	// if err != nil {
	// 	return fmt.Errorf("can't delete a bank.bic[%d]: error get list cards from db.card: %v", bic, err)
	// }

	// if len(listCard) != 0 {
	// 	return fmt.Errorf("can't delete a bank.bic[%d] if cards from this bank exist", bic)
	// }

	// return b.bankRepo.Delete(ctx, bic)
	return nil
}
