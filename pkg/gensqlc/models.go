// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package gensqlc

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID          uuid.UUID  `json:"id"`
	PassengerID *uuid.UUID `json:"passenger_id"`
	Type        string     `json:"type"`
	Number      string     `json:"number"`
}

type Passenger struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	MiddleName string    `json:"middle_name"`
}

type Ticket struct {
	ID              uuid.UUID `json:"id"`
	PassengerID     uuid.UUID `json:"passenger_id"`
	OrderNumber     string    `json:"order_number"`
	Departure       string    `json:"departure"`
	Destination     string    `json:"destination"`
	ServiceProvider string    `json:"service_provider"`
	DepartureDate   time.Time `json:"departure_date"`
	ArrivalDate     time.Time `json:"arrival_date"`
	OrderDate       time.Time `json:"order_date"`
}
