package dto

import (
	"encoding/json"
	"time"
)

type Password struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	UserID string `json:"user_id"`
	Value  string `json:"value"`
}

func (e *Password) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
