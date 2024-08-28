package database

type Card interface {
	CreateTable() (err error)
	Create(c Card) (err error)
	Read(id uint) (c Card, err error)
	List(c Card) (outc []Card, err error)
	Update(c Card) (err error)
	Delete(id uint) (err error)
}
