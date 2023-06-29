package shift

import (
	"github.com/Laura-2950/desafio-final-go/internal/domain"
	"github.com/Laura-2950/desafio-final-go/pkg/store"
	"github.com/Laura-2950/desafio-final-go/pkg/web"
)

type IRepository interface {
	//GetByID(id int) (*domain.Shift, error)
	CreateNewShift(shift *domain.Shift) (*domain.Shift, error)
	DeleteShift(id int) error
	//Update(shift *domain.Shift) (*domain.Shift, error)
}

type Repository struct {
	Storage store.StoreInterface
}

func (r *Repository) CreateNewShift(shift *domain.Shift) (*domain.Shift, error) {
	if r.Storage.Exists(shift.Dentist.RegistrationNumber, "registration_number", "dentists") {
		return nil, web.NewBadRequestApiError("existing dentist")
	}
	if r.Storage.Exists(shift.Patient.Dni, "dni", "patients") {
		return nil, web.NewBadRequestApiError("existing patient")
	}
	//falta valudidar el no exista el mismo turno a la misma hora y día para el mismo pasiente y/o dentista
	shift, err := r.Storage.CreateShift(*shift)
	if err != nil {
		return nil, web.NewInternalServerApiError("unexpected error")
	}
	return shift, nil
}