package dentist

import (
	"fmt"

	"github.com/Laura-2950/desafio-final-go/internal/domain"
	"github.com/Laura-2950/desafio-final-go/pkg/store"
	"github.com/Laura-2950/desafio-final-go/pkg/web"
)

type IRepository interface {
	GetByID(id int) (*domain.Dentist, error)
	CreateNewDentist(dentist *domain.Dentist) (*domain.Dentist, error)
}

type Repository struct {
	Storage store.StoreInterface
}

func (r *Repository) CreateNewDentist(dentist *domain.Dentist) (*domain.Dentist, error) {
	if r.Storage.Exist(dentist.RegistrationNumber) {
		return nil, web.NewBadRequestApiError("existing product")
	}
	dentist, err := r.Storage.CreateDentist(*dentist)
	if err != nil {
		return nil, web.NewInternalServerApiError("unexpected error")
	}
	return dentist, nil
}

func (r *Repository) GetByID(id int) (*domain.Dentist, error) {
	dent, err := r.Storage.ReadDentist(id)
	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("dentist_id %d not found", id))
	}
	return dent, nil
}
