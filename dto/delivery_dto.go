package dto

import (
	"encoding/json"
	"time"
)

var (
	NilDeliveryRequestType = DeliveryRequestType{}
	NilDeliveryRequest     = DeliveryRequest{}
	NilDeliveryAddress     = DeliveryAddress{}
)

type DeliveryRequestType struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	Name string `json:"name"`
}

func (e *DeliveryRequestType) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type DeliveryRequest struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	DeliveryAddressID string `json:"delivery_address_id"`

	DeliveryRequestTypeID string              `json:"delivery_request_type_id"`
	DeliveryRequestType   DeliveryRequestType `json:"delivery_request_type"`
	RequestMessage        string              `json:"request_message"`
}

func (e *DeliveryRequest) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type DeliveryAddress struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	UserID string `json:"user_id"`

	IsDefault       bool   `json:"is_default"`
	ReceiverName    string `json:"receiver_name"`
	ReceiverContact string `json:"receiver_contact"`
	PostCode        string `json:"post_code"`
	Address         string `json:"address"`
	AddressDetail   string `json:"address_detail"`

	DeliveryRequest DeliveryRequest `json:"delivery_request"`
}

func (e *DeliveryAddress) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
