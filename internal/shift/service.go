package shift

import "github.com/Laura-2950/desafio-final-go/internal/domain"

type IService interface {
	GetShiftByID(id int) (*domain.ResponseShift, error)
	CreateNewShift(shift *domain.Shift) (*domain.Shift, error)
	CreateNewShiftCode(shift *domain.ShiftCode) (*domain.Shift, error)
	Delete(id int) error
	UpdateShift(id int, shift *domain.Shift) (*domain.Shift, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) CreateNewShift(shift *domain.Shift) (*domain.Shift, error) {
	res, err := s.Repository.CreateNewShift(shift)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) CreateNewShiftCode(shift *domain.ShiftCode) (*domain.Shift, error) {
	res, err := s.Repository.CreateNewShiftCode(shift)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) Delete(id int) error {
	err := s.Repository.DeleteShift(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateShift(id int, shift *domain.Shift) (*domain.Shift, error) {
	sh, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if shift.Patient != 0 {
		sh.Patient = shift.Patient
	}
	if shift.Dentist != 0 {
		sh.Dentist = shift.Dentist
	}
	if shift.DateHour != "" {
		sh.DateHour = shift.DateHour
	}
	if shift.Description != "" {
		sh.Description = shift.Description
	}
	response, err := s.Repository.Update(sh)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Service) GetShiftByID(id int) (*domain.ResponseShift, error) {
	shift, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	response, err := s.Repository.TransformShiftToResponse(*shift)
	if err != nil {
		return nil, err
	}

	return response, nil
}
