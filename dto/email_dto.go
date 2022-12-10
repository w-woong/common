package dto

import "time"

var NilEmail = Email{}

type Email struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Priority uint8  `json:"priority"`
}

type Emails []Email
