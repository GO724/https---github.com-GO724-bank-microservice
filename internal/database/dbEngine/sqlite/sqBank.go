package sq

import (
	"bank-microservice/internal/entity"
	"context"
)

type bank struct {
	// pgInstance *postgres
	// err        error
}

func NewBank(ctx context.Context) (*bank, error) {
	b := bank{}
	// b.pgInstance, b.err = NewPG(ctx)
	return &b, nil
}

func (b bank) Migrate() error {
	// drop table

	// create table

	// fill table

	return nil
}

func (b bank) Create(bank entity.Bank) error {
	// New record
	return nil
}

func (b bank) Read(bic uint) (entity.Bank, error) {
	// Read record
	var bank entity.Bank
	bank.Bic = bic
	bank.Name = "name"
	return bank, nil
}

func (b bank) List(bank entity.Bank) ([]entity.Bank, error) {
	// List of record filtered on b(Bank)
	sb := make([]entity.Bank, 0, 10)
	return sb, nil
}

func (b bank) Update(bank entity.Bank) error {
	// Update record
	return nil
}

func (b bank) Delete(bic uint) error {
	// Delete record
	return nil
}
