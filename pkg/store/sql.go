package store

import (
	"database/sql"

	"github.com/Laura-2950/desafio-final-go/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) CreateDentist(dentist domain.Dentist) (*domain.Dentist, error) {
	query := "INSERT INTO dentists (name, last_name, registration_number) VALUES (?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return &domain.Dentist{}, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(dentist.Name, dentist.LastName, dentist.RegistrationNumber)
	if err != nil {
		return &domain.Dentist{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return &domain.Dentist{}, err
	}

	lid, _ := res.LastInsertId()
	dentist.ID = int(lid)
	return &dentist, nil
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
