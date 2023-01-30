package dto

import (
	"encoding/json"
	"time"
)

var NilUser = User{}

type User struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	LoginID     string `json:"login_id"`
	LoginType   string `json:"login_type"`
	LoginSource string `json:"login_source"`

	CredentialPassword *CredentialPassword `json:"password,omitempty"`
	CredentialToken    *CredentialToken    `json:"token,omitempty"`
	Personal           *Personal           `json:"personal,omitempty"`
	Emails             Emails              `json:"emails,omitempty"`
	DeliveryAddress    *DeliveryAddress    `json:"delivery_address,omitempty"`
	PaymentMethod      *PaymentMethod      `json:"payment_method,omitempty"`
}

func (d *User) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

type Users []User
