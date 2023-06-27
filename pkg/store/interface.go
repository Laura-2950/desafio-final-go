package store

import "github.com/Laura-2950/desafio-final-go/internal/domain"

type StoreInterface interface {
	ReadDentist(id int) (*domain.Dentist, error)
	CreateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	// UpdateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	Delete(id int, table string) error
	Exists(code, codeName, table string) bool
	ExistId(id int, table string) bool
	// ReadAllTurn() error
}
