package shift

import (
	"fmt"

	"github.com/Laura-2950/desafio-final-go/internal/domain"
	"github.com/Laura-2950/desafio-final-go/pkg/store"
	"github.com/Laura-2950/desafio-final-go/pkg/web"
)

type IRepository interface {
	//GetByID(id int) (*domain.Shift, error)
	CreateNewShift(shift *domain.Shift) (*domain.ResponseShift, error)
	DeleteShift(id int) error
	//Update(shift *domain.Shift) (*domain.Shift, error)
}

type Repository struct {
	Storage store.StoreInterface
}

func (r *Repository) CreateNewShift(shift *domain.Shift) (*domain.ResponseShift, error) {
	if r.Storage.ExistId(shift.Dentist, "dentists") {
		return nil, web.NewBadRequestApiError("nonexistent dentist")
	}
	if r.Storage.ExistId(shift.Patient, "patients") {
		return nil, web.NewBadRequestApiError("nonexistent patient")
	}
	if !r.Storage.ExistShift(shift)  {
		return nil, web.NewBadRequestApiError("existent shift")
	}

	res, err := r.Storage.CreateShift(shift)
	if err != nil {
		return nil, web.NewInternalServerApiError("unexpected error")
	}

	pat, err := r.Storage.ReadPatient(shift.Patient)
	if err != nil {
		return nil, err
	}
	dent, err := r.Storage.ReadDentist(shift.Dentist)
	if err != nil {
		return nil, err
	}	
	
	shiftCrete := domain.ResponseShift{
		ID:             res.ID,
		Patient:         domain.Patient{
			ID: pat.ID,
			Name: pat.Name,
			LastName: pat.LastName,
			Address: pat.Address,
			Dni: pat.Dni,
			RegistrationDate: pat.RegistrationDate,
		},
		Dentist:          domain.Dentist{
			ID: dent.ID,
			Name: dent.Name,
			LastName: dent.LastName,
			RegistrationNumber: dent.RegistrationNumber,
		},
		DateHour:              res.DateHour,
		Description: res.Description,
	}

	return &shiftCrete, nil
}

func (r *Repository) DeleteShift(id int) error {
	if r.Storage.ExistId(id, "shifts") {
		return web.NewBadRequestApiError(fmt.Sprintf("shift_id %d not found", id))
	}
	err := r.Storage.Delete(id, "shifts")
	if err != nil {
		return err
	}
	return nil
}
