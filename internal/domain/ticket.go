package domain

import (
	"github.com/google/uuid"
)

type Ticket struct {
	ID              uuid.UUID
	PassengerID     uuid.UUID
	OrderNumber     string
	Departure       string
	Destination     string
	ServiceProvider string
	DepartureDate   string
	ArrivalDate     string
	OrderDate       string
}
