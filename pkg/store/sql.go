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

func (s *SqlStore) Exists(code, codeName, table string) bool {
	var exist bool
	var id int

	query := "SELECT id FROM " + table + " WHERE " + codeName + " = ?;"
	row := s.DB.QueryRow(query, code)
	err := row.Scan(&id)
	if err != nil {
		return exist
	}

	if id > 0 {
		exist = true
	}

	return exist
}

func (s *SqlStore) ExistId(id int, table string) bool {
	var exist bool
	query := "SELECT * FROM " + table + " WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			exist = true
		}
	} else {
		exist = false
	}
	return exist
}
