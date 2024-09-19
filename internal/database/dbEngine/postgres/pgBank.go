package pg

import (
	"bank-microservice/internal/entity"
	"context"
	"fmt"
	"log"
	"os"
)

type bank struct {
	pgInstance *postgres
	err        error
}

func NewBank(ctx context.Context) (*bank, error) {
	b := bank{}
	b.pgInstance, b.err = NewPG(ctx)
	return &b, b.err
}

func (b bank) Migrate() error {
	// drop table
	err := b.pgInstance.ExecSql("DROP TABLE IF EXISTS public.bank")
	if err != nil {
		log.Fatal(err)
	}

	// create table
	err = b.pgInstance.ExecSql("CREATE TABLE IF NOT EXISTS public.bank (bic integer NOT NULL, name text NOT NULL, PRIMARY KEY (bic))")
	if err != nil {
		log.Fatal(err)
	}

	// fill table
	err = b.pgInstance.ExecSql("INSERT INTO public.bank (bic,name) VALUES ('123','CreditBank')")
	if err != nil {
		log.Fatal(err)
	}
	err = b.pgInstance.ExecSql("INSERT INTO public.bank (bic,name) VALUES ('456','DebitBank')")
	if err != nil {
		log.Fatal(err)
	}
	err = b.pgInstance.ExecSql("INSERT INTO public.bank (bic,name) VALUES ('789','SaldoBank')")
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (b bank) Create(bank entity.Bank) error {
	// New record
	res, err := b.pgInstance.db.Exec(*b.pgInstance.ctx, "INSERT INTO bank(bic,name) VALUES (%1,%2)", bank.Name, bank.Bic)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to create bank.bic[%v].name[%v] : %v\n", bank.Bic, bank.Name, err)
	} else {
		fmt.Fprintf(os.Stderr, "create bank.bic[%v].name[%v] : %v rows inserted\n", bank.Bic, bank.Name, res.RowsAffected())
	}
	return err
}

func (b bank) Read(bic uint) (entity.Bank, error) {
	// Read record
	sqlQuery := "SELECT bank.name FROM bank WHERE bank.bic=%1"

	var name string

	err := b.pgInstance.db.QueryRow(*b.pgInstance.ctx, sqlQuery, bic).Scan(&name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	var bank entity.Bank
	bank.Bic = bic
	bank.Name = name
	return bank, err
}

func (b bank) List(bank entity.Bank) ([]entity.Bank, error) {
	// List of record filtered on b(Bank)
	sb := make([]entity.Bank, 0, 10)
	return sb, nil
}

func (b bank) Update(bank entity.Bank) error {
	// Update record
	res, err := b.pgInstance.db.Exec(*b.pgInstance.ctx, "UPDATE bank SET name=%1 WHERE bic=%2", bank.Name, bank.Bic)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to update bank.bic[%v].name[%v] : %v\n", bank.Bic, bank.Name, err)
	} else {
		fmt.Fprintf(os.Stderr, "update bank.bic[%v].name[%v] : %v rows updated\n", bank.Bic, bank.Name, res.RowsAffected())
	}
	return err
}

func (b bank) Delete(bic uint) error {
	// Delete record
	res, err := b.pgInstance.db.Exec(*b.pgInstance.ctx, "DELETE FROM bank WHERE bic=%1", bic)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to delete bank.bic[%v]: %v\n", bic, err)
	} else {
		fmt.Fprintf(os.Stderr, "delete bank.bic[%v]: %v rows deleted\n", bic, res.RowsAffected())
	}
	return err
}
