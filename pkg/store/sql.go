package store

import (
	"database/sql"

	"github.com/Laura-2950/desafio-final-go/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) ReadDentist(id int) (*domain.Dentist, error) {
	var dentistReturn domain.Dentist

	query := "SELECT * FROM dentist WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&dentistReturn.ID, &dentistReturn.Name, &dentistReturn.LastName, &dentistReturn.RegistrationNumber)
	if err != nil {
		return nil, err
	}
	return &dentistReturn, nil
}
