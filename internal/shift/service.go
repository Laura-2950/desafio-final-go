package shift

import "github.com/Laura-2950/desafio-final-go/internal/domain"

type IService interface {
	//GetShiftByID(id int) (*domain.Shift, error)
	CreateNewShift(shift *domain.Shift) (*domain.ResponseShift, error)
	Delete(id int) error
	//UpdateShift(id int, dent *domain.Shift) (*domain.Shift, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) CreateNewShift(shift *domain.Shift) (*domain.ResponseShift, error) {
	res, err := s.Repository.CreateNewShift(shift)
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
