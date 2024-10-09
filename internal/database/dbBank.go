package database

import "bank-microservice/internal/entity"

type Bank interface {
	Migrate() error                            // Create table & fill demo record
	Create(b entity.Bank) error                // New record
	Read(bic uint) (entity.Bank, error)        // Read record
	List(b entity.Bank) ([]entity.Bank, error) // List of record filtered on b(Bank)
	Update(b entity.Bank) error                // Update record
	Delete(bic uint) error                     // Delete record
}
