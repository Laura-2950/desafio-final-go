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
	DeleteDentist(id int) error
	Update(dent *domain.Dentist) (*domain.Dentist, error)
}

type Repository struct {
	Storage store.StoreInterface
}

func (r *Repository) CreateNewDentist(dentist *domain.Dentist) (*domain.Dentist, error) {
	if r.Storage.Exists(dentist.RegistrationNumber, "registration_number", "dentists") {
		return nil, web.NewBadRequestApiError("existing dentist")
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

func (r *Repository) DeleteDentist(id int) error {
	if r.Storage.ExistId(id, "dentists") {
		return web.NewBadRequestApiError(fmt.Sprintf("nonexistent dentist with id %d.", id))
	}
	err := r.Storage.Delete(id, "dentists")
	if err != nil {
		return web.NewNotFoundApiError(fmt.Sprintf("dentist_id %d not found", id))
	}
	return nil
}

// PUT y PATCH ???
func (r *Repository) Update(dent *domain.Dentist) (*domain.Dentist, error) {
	// if r.Storage.Exists(dent.RegistrationNumber, "registration_name", "dentists") {
	// 	return nil, web.NewExistsError(fmt.Sprintf("dentist_registrationNumber"))
	// }
	dentist, err := r.Storage.UpdateDentist(*dent)
	if err != nil {
		return nil, web.NewUpdateError(fmt.Sprintf("error updating dentist"))
	}
	return dentist, nil
}
