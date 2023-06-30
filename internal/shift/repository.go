package shift

import (
	"fmt"

	"github.com/Laura-2950/desafio-final-go/internal/domain"
	"github.com/Laura-2950/desafio-final-go/pkg/store"
	"github.com/Laura-2950/desafio-final-go/pkg/web"
)

type IRepository interface {
	GetByID(id int) (*domain.Shift, error)
	CreateNewShift(shift *domain.Shift) (*domain.Shift, error)
	CreateNewShiftCode(shift *domain.ShiftCode) (*domain.Shift, error)
	DeleteShift(id int) error
	Update(shift *domain.Shift) (*domain.Shift, error)
	TransformShiftToResponse(shift domain.Shift) (*domain.ResponseShift, error)
}

type Repository struct {
	Storage store.StoreInterface
}

func (r *Repository) CreateNewShift(shift *domain.Shift) (*domain.Shift, error) {
	if r.Storage.ExistId(shift.Dentist, "dentists") {
		return nil, web.NewBadRequestApiError("nonexistent dentist")
	}
	if r.Storage.ExistId(shift.Patient, "patients") {
		return nil, web.NewBadRequestApiError("nonexistent patient")
	}
	if !r.Storage.ExistShift(shift) {
		return nil, web.NewBadRequestApiError("existent shift")
	}

	shiftCreate, err := r.Storage.CreateShift(shift)
	if err != nil {
		return nil, web.NewInternalServerApiError("unexpected error")
	}

	return shiftCreate, nil
}

func (r *Repository) CreateNewShiftCode(shift *domain.ShiftCode) (*domain.Shift, error) {
	patient, err := r.Storage.ReadPatientByDNI(shift.Patient)
	if err != nil {
		return nil, err
	}
	dentist, err := r.Storage.ReadDentistByCode(shift.Dentist)
	if err != nil {
		return nil, err
	}
	aux := domain.Shift{
		Patient:     patient.ID,
		Dentist:     dentist.ID,
		DateHour:    shift.DateHour,
		Description: shift.Description,
	}
	if !r.Storage.ExistShift(&aux) {
		return nil, web.NewBadRequestApiError("existent shift")
	}
	shiftCreate, err := r.Storage.CreateShift(&aux)
	if err != nil {
		return nil, web.NewInternalServerApiError("unexpected error")
	}
	return shiftCreate, nil
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

func (r *Repository) Update(shift *domain.Shift) (*domain.Shift, error) {
	response, err := r.Storage.UpdateShift(*shift)
	if err != nil {
		return nil, web.NewInternalServerApiError("error updating shift")
	}

	return response, nil
}

func (r *Repository) GetByID(id int) (*domain.Shift, error) {
	response, err := r.Storage.ReadShift(id)
	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("shift_id %d not found", id))
	}

	return response, nil
}

func (s *Repository) TransformShiftToResponse(shift domain.Shift) (*domain.ResponseShift, error) {
	pat, err := s.Storage.ReadPatient(shift.Patient)
	if err != nil {
		return nil, err
	}
	dent, err := s.Storage.ReadDentist(shift.Dentist)
	if err != nil {
		return nil, err
	}

	shiftCreate := domain.ResponseShift{
		ID: shift.ID,
		Patient: domain.Patient{
			ID:               pat.ID,
			Name:             pat.Name,
			LastName:         pat.LastName,
			Address:          pat.Address,
			Dni:              pat.Dni,
			RegistrationDate: pat.RegistrationDate,
		},
		Dentist: domain.Dentist{
			ID:                 dent.ID,
			Name:               dent.Name,
			LastName:           dent.LastName,
			RegistrationNumber: dent.RegistrationNumber,
		},
		DateHour:    shift.DateHour,
		Description: shift.Description,
	}

	return &shiftCreate, nil
}
