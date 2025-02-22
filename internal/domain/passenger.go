package domain

import (
	"github.com/google/uuid"
)

type Passenger struct {
	ID         uuid.UUID
	FirstName  string
	LastName   string
	MiddleName string
	Documents  []Document
	Tickets    []Ticket
}
