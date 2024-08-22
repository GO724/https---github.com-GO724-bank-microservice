package postgrees

type Card interface {
	CreateTableCard() (err error)
	CreateCard(c Card) (err error)
	ReadCard(id uint) (c Card, err error)
	UpdateCard(c Card) (err error)
	DeleteCard(id uint) (err error)
}
