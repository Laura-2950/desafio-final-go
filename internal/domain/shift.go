package domain

type Shift struct {
	ID                 int    `json:"id"`
	Patient            Patient `json:"patient" binding:"required"`
	Dentist            Dentist `json:"dentist" binding:"required"`
	DateHour           string `json:"date_hour" binding:"required"`
}

type RequestShift struct {
	Patient            Patient `json:"patient" binding:"required"`
	Dentist            Dentist `json:"dentist" binding:"required"`
	DateHour           string `json:"date_hour" binding:"required"`
}