package patient

import (
	"fmt"

	"github.com/Laura-2950/desafio-final-go/internal/domain"
	"github.com/Laura-2950/desafio-final-go/pkg/store"
	"github.com/Laura-2950/desafio-final-go/pkg/web"
)

type IRepository interface {
	GetByID(id int) (*domain.Patient, error)
	//CreateNewPatient(patient *domain.Patient) (*domain.Patient, error)
	//DeletePatient(id int) error
	Update(pat *domain.Patient) (*domain.Patient, error)
}

type Repository struct {
	Storage store.StoreInterface
}

func (r *Repository) GetByID(id int) (*domain.Patient, error) {
	patient, err := r.Storage.ReadPatient(id)
	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("patient_id %d not found", id))
	}
	return patient, nil
}

// PUT y PATCH
func (r *Repository) Update(pat *domain.Patient) (*domain.Patient, error) {
	patient, err := r.Storage.UpdatePatient(*pat)
	if err != nil {
		return nil, web.NewInternalServerApiError("error updating patient")
	}
	return patient, nil
}