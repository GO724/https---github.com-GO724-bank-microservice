package bank

import "context"

func (b *bankService) New(ctx context.Context, bic uint) error {
	// bank, err := b.bankRepo.Read(ctx, bic)
	// if err != nil {
	// 	return fmt.Errorf("can't create a bank.bic[%d]: read error: %v", bic, err)
	// }

	// if

	// // some logic: can't delete a bank if cards from this bank exist

	// searchCard := entity.Card{Id: 0,
	// 	Person:  entity.Person{Inn: 0, Name: ""},
	// 	Bank:    bank,
	// 	Expires: time.Date(10000, time.January, 1, 0, 0, 0, 0, &time.Location{}),
	// }

	// listCard, err := cardrepo.Card.List(ctx, searchCard)
	// if err != nil {
	// 	return fmt.Errorf("can't delete a bank.bic[%d]: error get list cards from db.card: %v", bic, err)
	// }

	// if len(listCard) != 0 {
	// 	return fmt.Errorf("can't delete a bank.bic[%d] if cards from this bank exist", bic)
	// }

	// return b.bankRepo.Delete(ctx, bic)
	return nil
}
