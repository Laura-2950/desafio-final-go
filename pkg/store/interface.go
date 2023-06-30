package store

import "github.com/Laura-2950/desafio-final-go/internal/domain"

type StoreInterface interface {
	ReadDentist(id int) (*domain.Dentist, error)
	CreateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	UpdateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	Delete(id int, table string) error
	Exists(code, codeName, table string) bool
	ExistId(id int, table string) bool

	ReadPatient(id int) (*domain.Patient, error)
	CreatePatient(patient domain.Patient) (*domain.Patient, error)
	UpdatePatient(patient domain.Patient) (*domain.Patient, error)

	CreateShift(shift *domain.Shift) (*domain.Shift, error)
	ExistShift(shift *domain.Shift) bool
	UpdateShift(shift domain.Shift) (*domain.Shift, error)
	ReadShift(id int) (*domain.Shift, error)
	// ReadAllTurn() error
}
