package dto

import (
	"encoding/json"
	"time"
)

var (
	NilCredentialPassword = CredentialPassword{}
)

type CredentialPassword struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	UserID string `json:"user_id"`
	Value  string `json:"value"`
}

func (e *CredentialPassword) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type CredentialToken struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	UserID string `json:"user_id"`
	Value  string `json:"value"`
}

func (e *CredentialToken) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
