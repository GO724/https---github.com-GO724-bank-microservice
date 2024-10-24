package card

import (
	"bank-microservice/internal/entity"
	"context"
)

func (c cardRepository) Set(ctx context.Context, card entity.Card) error {
	// Update record
	//	res, err := b.db.db.Exec(*b.db.ctx, "UPDATE bank SET name=%1 WHERE bic=%2", bank.Name, bank.Bic)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "unable to update bank.bic[%v].name[%v] : %v\n", bank.Bic, bank.Name, err)
	//	} else {
	//		fmt.Fprintf(os.Stderr, "update bank.bic[%v].name[%v] : %v rows updated\n", bank.Bic, bank.Name, res.RowsAffected())
	//	}
	//	return err
	return nil
}
