package patient

import "github.com/Laura-2950/desafio-final-go/internal/domain"

type IService interface {
	GetPatientByID(id int) (*domain.Patient, error)
	CreateNewPatient(patient *domain.Patient) (*domain.Patient, error)
	DeletePatient(id int) error
	UpdatePatient(id int, dent *domain.Patient) (*domain.Patient, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) GetPatientByID(id int) (*domain.Patient, error) {
	patient, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return patient, nil
}

func (s *Service) UpdatePatient(id int, pat *domain.Patient) (*domain.Patient, error) {
	patient, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	if pat.Name != "" {
		patient.Name = pat.Name
	}
	if pat.LastName != "" {
		patient.LastName = pat.LastName
	}
	if pat.Address != "" {
		patient.Address = pat.Address
	}
	patient, err = s.Repository.Update(patient)
	if err != nil {
		return nil, err
	}
	return patient, nil
}

func (s *Service) CreateNewPatient(patient *domain.Patient) (*domain.Patient, error) {
	p, err := s.Repository.CreateNewPatient(patient)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Service) DeletePatient(id int) error {
	err := s.Repository.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
}
