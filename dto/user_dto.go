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

	CredentialPassword CredentialPassword `json:"password"`
	CredentialToken    CredentialToken    `json:"token"`
	Personal           Personal           `json:"personal"`
	Emails             Emails             `json:"emails"`
	DeliveryAddress    DeliveryAddress    `json:"delivery_address"`
	PaymentMethod      PaymentMethod      `json:"payment_method"`
}

func (d *User) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

type Users []User
