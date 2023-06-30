package domain

type Shift struct {
	ID                 int    `json:"id"`
	Patient            int `json:"patient" binding:"required"`
	Dentist            int `json:"dentist" binding:"required"`
	DateHour           string `json:"date_hour" binding:"required"`
	Description        string `json:"description"`
}

type ShiftCode struct {
	ID                 int    `json:"id"`
	Patient            string `json:"patient_dni" binding:"required"`
	Dentist            string `json:"dentist_registration_number" binding:"required"`
	DateHour           string `json:"date_hour" binding:"required"`
	Description        string `json:"description"`
}

type ResponseShift struct {
	ID                 int    `json:"id"`
	Patient            Patient `json:"patient" binding:"required"`
	Dentist            Dentist `json:"dentist" binding:"required"`
	DateHour           string `json:"date_hour" binding:"required"`
	Description        string `json:"description"`
}

type RequestShift struct {
	Patient            int `json:"patient,omitempty"`
	Dentist            int `json:"dentist,omitempty"`
	DateHour           string `json:"date_hour,omitempty"`
	Description        string `json:"description,omitempty"`
}