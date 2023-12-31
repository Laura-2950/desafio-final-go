package dentist

import "github.com/Laura-2950/desafio-final-go/internal/domain"

type IService interface {
	GetDentistByID(id int) (*domain.Dentist, error)
	CreateNewDentist(dentist *domain.Dentist) (*domain.Dentist, error)
	DeleteDentist(id int) error
	UpdateDentist(id int, dent *domain.Dentist) (*domain.Dentist, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) CreateNewDentist(dentist *domain.Dentist) (*domain.Dentist, error) {
	dent, err := s.Repository.CreateNewDentist(dentist)
	if err != nil {
		return nil, err
	}
	return dent, nil
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

func (s *Service) UpdateDentist(id int, dent *domain.Dentist) (*domain.Dentist, error) {
	dentist, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	if dent.Name != "" {
		dentist.Name = dent.Name
	}
	if dent.LastName != "" {
		dentist.LastName = dent.LastName
	}

	dentist, err = s.Repository.Update(dentist)
	if err != nil {
		return nil, err
	}
	return dentist, nil
}
