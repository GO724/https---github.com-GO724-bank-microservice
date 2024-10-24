package card

import (
	"context"
)

func (c cardRepository) Migrate(ctx context.Context) error {
	// // drop table
	// err := c.db.ExecSql(ctx, "DROP TABLE IF EXISTS public.card")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // create table
	// err = c.db.ExecSql(ctx, "CREATE TABLE IF NOT EXISTS public.card (id integer NOT NULL, person_inn integer NOT NULL, bank_bic integer NOT NULL, expiries date NOT NULL, PRIMARY KEY (id)))")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = c.db.ExecSql(ctx, "ALTER TABLE IF EXISTS public.card ADD FOREIGN KEY (person_inn) REFERENCES public.person (inn) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = c.db.ExecSql(ctx, "ALTER TABLE IF EXISTS public.card ADD FOREIGN KEY (bank_bic) REFERENCES public.bank (bic) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // fill table
	// err = c.db.ExecSql(ctx, "INSERT INTO public.card (id,bic,name) VALUES ('123','CreditBank')")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = c.db.ExecSql(ctx, "INSERT INTO public.bank (bic,name) VALUES ('456','DebitBank')")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = c.db.ExecSql(ctx, "INSERT INTO public.bank (bic,name) VALUES ('789','SaldoBank')")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// return err
	return nil
}
