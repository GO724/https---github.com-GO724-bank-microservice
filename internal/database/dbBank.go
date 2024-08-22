package database

type Bank interface {
	CreateTable() (err error)
	CreateBank(b Bank) (err error)
	ReadBank(bic uint) (outb Bank, err error)
	UpdateBank(b Bank) (err error)
	DeleteBank(bic uint) (err error)
}
