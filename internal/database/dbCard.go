package database

import "bank-microservice/internal/entity"

type Card interface {
	Migrate() error                            // Create table & fill demo record
	Create(c entity.Card) error                // New record
	Read(id uint) (entity.Card, error)         // Read record
	List(c entity.Card) ([]entity.Card, error) // List of record filtered on c(Card)
	Update(c entity.Card) error                // Update record
	Delete(id uint) error                      // Delete record
	IsValid(c entity.Card) bool                // Check valid (c.Expires.Before(time.Now().UTC()))
}
