package person

import (
	"context"
	"fmt"
)

func (p personRepository) Delete(ctx context.Context, inn uint) error {
	// // Delete record
	// err := p.db.ExecSql(ctx, fmt.Sprintf("DELETE FROM person WHERE inn=%d", inn))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "unable to delete person.inn[%d] : %v\n", inn, err)
	// } else {
	// 	fmt.Fprintf(os.Stderr, "delete person.inn[%d]\n", inn)
	// }
	// return err
	fmt.Println("service: delete person.inn[%d]", inn)
	return nil
}
