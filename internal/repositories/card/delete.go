package card

import (
	"context"
)

func (c *cardRepository) Delete(ctx context.Context, id uint) error {
	// // Delete record
	// err := c.db.ExecSql(ctx, fmt.Sprintf("DELETE FROM card WHERE id=%d", id))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "unable to delete card.id[%d] : %v\n", id, err)
	// } else {
	// 	fmt.Fprintf(os.Stderr, "delete card.id[%d]\n", id)
	// }
	// return err
	return nil
}
