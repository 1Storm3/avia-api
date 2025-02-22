package model

import (
	"github.com/google/uuid"
)

type Ticket struct {
	ID              uuid.UUID `db:"id"`
	PassengerID     uuid.UUID `db:"passenger_id"`
	OrderNumber     string    `db:"order_number"`
	Departure       string    `db:"departure"`
	Destination     string    `db:"destination"`
	ServiceProvider string    `db:"service_provider"`
	DepartureDate   string    `db:"departure_date"`
	ArrivalDate     string    `db:"arrival_date"`
	OrderDate       string    `db:"order_date"`
}
