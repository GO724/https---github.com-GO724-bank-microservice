package bank

import (
	"bank-microservice/internal/entity"
	"context"
)

func (b bankRepository) Set(ctx context.Context, bank entity.Bank) error {
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
