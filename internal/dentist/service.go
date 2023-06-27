package dentist

import "github.com/Laura-2950/desafio-final-go/internal/domain"

type IService interface {
	GetDentistByID(id int) (*domain.Dentist, error)
	CreateNewDentist(dentist *domain.Dentist) (*domain.Dentist, error)
	DeleteDentist(id int) error
}

type Service struct {
	Repository IRepository
}

func (s *Service) CreateNewDentist(dentist *domain.Dentist) (*domain.Dentist, error) {
	product, err := s.Repository.CreateNewDentist(dentist)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *Service) GetDentistByID(id int) (*domain.Dentist, error) {
	dentist, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return dentist, nil
}

func (s *Service) DeleteDentist(id int) error {
	err := s.Repository.DeleteDentist(id)
	if err != nil {
		return err
	}
	return nil
}
