package database

type Bank interface {
	CreateTable() (err error)
	Create(b Bank) (err error)
	Read(bic uint) (outb Bank, err error)
	List(b Bank) (outb []Bank, err error)
	Update(b Bank) (err error)
	Delete(bic uint) (err error)
}
