package database

type Person interface {
	CreateTable() (err error)
	CreatePerson(p Person) (err error)
	ReadPerson(inn uint) (p Person, err error)
	UpdatePerson(p Person) (err error)
	DeletePerson(inn uint) (err error)
}
