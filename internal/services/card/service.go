package card

import (
	"bank-microservice/internal/entity"
	repo "bank-microservice/internal/repositories/card"
	"context"
)

type CardRepo interface {
	New(ctx context.Context, c entity.Card) (*entity.Card, error)    // New record
	Get(ctx context.Context, id uint) (*entity.Card, error)          // Read record
	Set(ctx context.Context, c entity.Card) (*entity.Card, error)    // Update record
	Delete(ctx context.Context, id uint) error                       // Delete record
	List(ctx context.Context, c entity.Card) ([]*entity.Card, error) // List of record filtered on c(Card)
}

type cardService struct {
	cardRepo repo.Card
}

func NewService(rc repo.Card) *cardService {
	return &cardService{
		cardRepo: rc,
	}
}
