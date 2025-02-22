package model

import (
	"github.com/google/uuid"
)

type Passenger struct {
	ID         uuid.UUID  `db:"id"`
	FirstName  string     `db:"first_name"`
	LastName   string     `db:"last_name"`
	MiddleName string     `db:"middle_name"`
	Documents  []Document `db:"documents"`
	Tickets    []Ticket   `db:"tickets"`
}
