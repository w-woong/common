package dto

import (
	"encoding/json"
	"time"
)

var (
	NilPaymentType   = PaymentType{}
	NilPaymentMethod = PaymentMethod{}
)

type PaymentType struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	Name string `json:"name"`
}

func (e *PaymentType) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type PaymentMethod struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	UserID string `json:"user_id"`

	PaymentTypeID string      `json:"payment_type_id"`
	PaymentType   PaymentType `json:"payment_type"`
	Identity      string      `json:"identity"`
	Option        string      `json:"option"`
}

func (e *PaymentMethod) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
