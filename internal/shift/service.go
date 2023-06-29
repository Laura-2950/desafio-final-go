package shift

import "github.com/Laura-2950/desafio-final-go/internal/domain"

type IService interface {
	//GetShiftByID(id int) (*domain.Shift, error)
	CreateNewShift(shift *domain.Shift) (*domain.Shift, error)
	//DeleteShift(id int) error
	//UpdateShift(id int, dent *domain.Shift) (*domain.Shift, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) CreateNewShift(shift *domain.Shift) (*domain.Shift, error) {
	shift, err := s.Repository.CreateNewShift(shift)
	if err != nil {
		return nil, err
	}
	return shift, nil
}
