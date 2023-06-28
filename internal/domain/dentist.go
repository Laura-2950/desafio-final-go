package domain

type Dentist struct {
	ID                 int    `json:"id"`
	Name               string `json:"name" binding:"required"`
	LastName           string `json:"last_name" binding:"required"`
	RegistrationNumber string `json:"registration_number" binding:"required"`
}

type RequestDentist struct {
	Name               string `json:"name,omitempty"`
	LastName           string `json:"last_name,omitempty"`
	RegistrationNumber string `json:"registration_number,omitempty"`
}
