package person

import (
	"context"
)

func (p personRepository) Migrate(ctx context.Context) error {
	// // drop table
	// err := p.db.ExecSql(ctx, "DROP TABLE IF EXISTS public.person")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // create table
	// err = p.db.ExecSql(ctx, "CREATE TABLE IF NOT EXISTS public.person (inn integer NOT NULL, name text NOT NULL, PRIMARY KEY (inn))")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // fill table
	// err = p.db.ExecSql(ctx, "INSERT INTO public.person (inn,name) VALUES ('123','Ivanov')")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = p.db.ExecSql(ctx, "INSERT INTO public.person (inn,name) VALUES ('456','Petrov')")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = p.db.ExecSql(ctx, "INSERT INTO public.person (inn,name) VALUES ('789','Sidorov')")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// return err
	return nil
}
