package database

type Person interface {
	CreateTable() (err error)
	Create(p Person) (err error)
	Read(inn uint) (p Person, err error)
	List(p Person) (outp []Person, err error)
	Update(p Person) (err error)
	Delete(inn uint) (err error)
}
