package postgrees

import (
	entity "bdms/internal/entity"
)

type pgbank struct {
}

func (b pgbank) CreateTableBank() (err error) {
	// create new table in new database
	return err
}

func (b pgbank) CreateBank(B entity.Bank) (err error) {
	// create new bank item (record in db)
	return err
}

func (b pgbank) ReadBank(bic uint) (outb entity.Bank, err error) {
	// seek record (PK bic) & return struct entity.Bank
	return outb, err
}

func (b pgbank) UpdateBank(inb entity.Bank) (err error) {
	// update record (PK inb.bic struct entity.Bank)
	return err
}

func (b pgbank) DeleteBank(bic uint) (err error) {
	// delete record (PK bic)
	return err
}
