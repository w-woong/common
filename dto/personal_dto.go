package dto

import "time"

type Personal struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	UserID string `json:"user_id"`

	FirstName   string     `json:"first_name"`
	MiddleName  string     `json:"middle_name"`
	LastName    string     `json:"last_name"`
	BirthYear   int        `json:"birth_year"`
	BirthMonth  int        `json:"birth_month"`
	BirthDay    int        `json:"birth_day"`
	BirthDate   *time.Time `json:"birth_date"`
	Gender      string     `json:"gender"`
	Nationality string     `json:"nationality"`
}
