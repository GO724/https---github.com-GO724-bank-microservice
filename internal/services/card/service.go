package card

import cardrepo "bank-microservice/internal/repositories/card"

type Card interface {
}

type cardService struct {
	cardRepo cardrepo.Card
}

func NewService(c cardrepo.Card) *cardService {
	return &cardService{
		cardRepo: c,
	}
}
