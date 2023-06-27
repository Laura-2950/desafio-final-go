package store

import "github.com/Laura-2950/desafio-final-go/internal/domain"

type StoreInterface interface {
	ReadDentist(id int) (*domain.Dentist, error)
	CreateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	// UpdateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	// Delete(id int) error
	// Exist(codeValue string) bool
	// ExistId(id int) bool
	// ReadAllTurn() error
}
