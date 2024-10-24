package bank

import (
	"bank-microservice/internal/entity"
	"context"
)

func (b bankRepository) New(ctx context.Context, bank entity.Bank) error {
	// New record
	// rowsAffected, err := b.db.ExecSql(ctx, fmt.Sprintf("INSERT INTO bank(bic,name) VALUES (%s,%d)", bank.Name, bank.Bic))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "unable to create bank.bic[%v].name[%v] : %v\n", bank.Bic, bank.Name, err)
	// } else {
	// 	fmt.Fprintf(os.Stderr, "create bank.bic[%v].name[%v]\n", bank.Bic, bank.Name)
	// }
	//return err
	return nil
}
