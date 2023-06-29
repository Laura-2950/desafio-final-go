package domain

type Patient struct {
	ID                 int    `json:"id"`
	Name               string `json:"name" binding:"required"`
	LastName           string `json:"last_name" binding:"required"`
	Address            string `json:"address"`
	Dni                string `json:"dni" binding:"required"`
	RegistrationDate   string `json:"registration_date" binding:"required"`
}

//nombre, apellido, domicilio, DNI y fecha de alta.

type RequestPatient struct {
	Name               string `json:"name,omitempty"`
	LastName           string `json:"last_name,omitempty"`
	Address            string `json:"address,omitempty"`
	Dni                string `json:"dni,omitempty"`
	RegistrationDate   string `json:"registration_date,omitempty"`
}