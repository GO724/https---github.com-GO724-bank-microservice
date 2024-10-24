package bank

import (
	"context"
)

func (b *bankRepository) Delete(ctx context.Context, bic uint) error {
	// // Delete record
	// err := b.db.ExecSql(ctx, fmt.Sprintf("DELETE FROM bank WHERE bic=%d", bic))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "unable to delete bank.bic[%d] : %v\n", bic, err)
	// } else {
	// 	fmt.Fprintf(os.Stderr, "delete bank.bic[%d]\n", bic)
	// }
	// return err
	return nil
}
