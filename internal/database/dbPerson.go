package database

import "bank-microservice/internal/entity"

type Person interface {
	Migrate() error                                // Create table & fill demo record
	Create(p entity.Person) error                  // New record
	Read(inn uint) (entity.Person, error)          // Read record
	List(p entity.Person) ([]entity.Person, error) // List of record filtered on p(Person)
	Update(p entity.Person) error                  // Update record
	Delete(inn uint) error                         // Delete record
}
