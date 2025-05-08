package entity

import (
	"time"
)

type OrderStatus string

const (
	Unknown    OrderStatus = ""
	Pending    OrderStatus = "Pendente"
	InProgress OrderStatus = "Preparando"
	Completed  OrderStatus = "Pronto"
	Delivering OrderStatus = "A caminho"
	Delivered  OrderStatus = "Entregue"
	Cancelled  OrderStatus = "Cancelado"
)

func (s OrderStatus) String() string {
	switch s {
	case Unknown:
		return ""
	case Pending:
		return "Pendente"
	case InProgress:
		return "Preparando"
	case Completed:
		return "Pronto"
	case Delivering:
		return "A caminho"
	case Delivered:
		return "Entregue"
	case Cancelled:
		return "Cancelado"
	}
	return ""
}

type Order struct {
	ID          ID
	Client      string
	Item        string
	Quantity    int
	Observation *string
	Status      OrderStatus
	CreatedAt   time.Time
}

func NewOrder(
	client string,
	item string,
	quantity int,
	observation *string,
	status OrderStatus,
) (*Order, error) {
	return &Order{
		ID:          NewID(),
		Client:      client,
		Item:        item,
		Quantity:    quantity,
		Observation: observation,
		Status:      status,
		CreatedAt:   time.Now().UTC(),
	}, nil
}
