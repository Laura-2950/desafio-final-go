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

	query := "SELECT * FROM dentists WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&dentistReturn.ID, &dentistReturn.Name, &dentistReturn.LastName, &dentistReturn.RegistrationNumber)
	if err != nil {
		return nil, err
	}
	return &dentistReturn, nil
}

func (s *SqlStore) Delete(id int, table string) error {
	stmt := "DELETE FROM " + table + " WHERE id = ?"
	_, err := s.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

// PUT & PATCH
func (s *SqlStore) UpdateDentist(dentist domain.Dentist) (*domain.Dentist, error) {
	query := "UPDATE dentists SET name=?, last_name=?, registration_number=?  WHERE id=?;"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(&dentist.Name, &dentist.LastName, &dentist.RegistrationNumber, &dentist.ID)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &dentist, nil
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

//-------------------------------Patient-------------------------------//

func (s *SqlStore) ReadPatient(id int) (*domain.Patient, error) {
	var patientReturn domain.Patient

	query := "SELECT * FROM patients WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&patientReturn.ID, &patientReturn.Name, &patientReturn.LastName, &patientReturn.Address, &patientReturn.Dni, &patientReturn.RegistrationDate)
	if err != nil {
		return nil, err
	}
	return &patientReturn, nil
}

func (s *SqlStore) ReadPatientByDNI(dni string) (*domain.Patient, error) {
	var patientReturn domain.Patient

	query := "SELECT * FROM patients WHERE dni = ?;"
	row := s.DB.QueryRow(query, dni)
	err := row.Scan(&patientReturn.ID, &patientReturn.Name, &patientReturn.LastName, &patientReturn.Address, &patientReturn.Dni, &patientReturn.RegistrationDate)
	if err != nil {
		return nil, err
	}
	return &patientReturn, nil
}

// PUT & PATCH
func (s *SqlStore) UpdatePatient(patient domain.Patient) (*domain.Patient, error) {
	query := "UPDATE patients SET name=?, last_name=?, address=?, dni=?, registration_date=? WHERE id=?;"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(&patient.Name, &patient.LastName, &patient.Address, &patient.Dni, &patient.RegistrationDate, &patient.ID)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &patient, nil
}

func (s *SqlStore) CreatePatient(patient domain.Patient) (*domain.Patient, error) {
	query := "INSERT INTO patients (name, last_name, address, dni, registration_date) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return &domain.Patient{}, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(patient.Name, patient.LastName, patient.Address, patient.Dni, patient.RegistrationDate)
	if err != nil {
		return &domain.Patient{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return &domain.Patient{}, err
	}

	lid, _ := res.LastInsertId()
	patient.ID = int(lid)

	return &patient, nil
}

//--------------------------Shift-----------------------------------------------------//

func (s *SqlStore) CreateShift(shift *domain.Shift) (*domain.Shift, error) {
	query := "INSERT INTO shifts (patient_id, dentist_id, date_hour, description) VALUES (?, ?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return &domain.Shift{}, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(shift.Patient, shift.Dentist, shift.DateHour, shift.Description)
	if err != nil {
		return &domain.Shift{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return &domain.Shift{}, err
	}

	lid, _ := res.LastInsertId()
	shift.ID = int(lid)

	return shift, nil
}

func (s *SqlStore) ExistShift(shift *domain.Shift) bool {
	var exist bool
	var id int
	query := "SELECT * FROM shifts WHERE patient_id = ? and dentist_id = ? and date_hour = ?;"
	row := s.DB.QueryRow(query, shift.Patient, shift.Dentist, shift.DateHour)
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
func (s *SqlStore) UpdateShift(shift domain.Shift) (*domain.Shift, error) {
	query := "UPDATE shifts SET patient_id=?, dentist_id=?, date_hour=?, description=? WHERE id=?;"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(&shift.Patient, &shift.Dentist, &shift.DateHour, &shift.Description, &shift.ID)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &shift, nil
}

func (s *SqlStore) ReadShift(id int) (*domain.Shift, error) {
	var shiftReturn domain.Shift

	query := "SELECT * FROM shifts WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&shiftReturn.ID, &shiftReturn.Patient, &shiftReturn.Dentist, &shiftReturn.DateHour, &shiftReturn.Description)
	if err != nil {
		return nil, err
	}
	return &shiftReturn, nil
}
