package bank

import (
	"context"
)

func (b bankRepository) Migrate(ctx context.Context) error {
	// // drop table
	// err := b.db.ExecSql(ctx, "DROP TABLE IF EXISTS public.bank")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // create table
	// err = b.db.ExecSql(ctx, "CREATE TABLE IF NOT EXISTS public.bank (bic integer NOT NULL, name text NOT NULL, PRIMARY KEY (bic))")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // fill table
	// err = b.db.ExecSql(ctx, "INSERT INTO public.bank (bic,name) VALUES ('123','CreditBank')")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = b.db.ExecSql(ctx, "INSERT INTO public.bank (bic,name) VALUES ('456','DebitBank')")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = b.db.ExecSql(ctx, "INSERT INTO public.bank (bic,name) VALUES ('789','SaldoBank')")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// return err
	return nil
}
