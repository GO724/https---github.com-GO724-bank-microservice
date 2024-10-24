package card

import (
	"bank-microservice/internal/entity"
	"context"
)

func (c cardRepository) New(ctx context.Context, card entity.Card) error {
	// // New record
	// err := c.db.ExecSql(ctx, fmt.Sprintf("INSERT INTO card(id,person_inn,bank_bic,expiries) VALUES (%d,%d,%d,%v)", card.Id, card.Bank.Bic, card.Person.Inn, card.Expires))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "unable to create card[%v] : %v\n", card, err)
	// } else {
	// 	fmt.Fprintf(os.Stderr, "create card[%v]\n", card)
	// }
	// return err
	return nil
}
