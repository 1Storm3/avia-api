package model

import (
	"github.com/google/uuid"
)

type Document struct {
	ID          uuid.UUID `db:"id"`
	PassengerID uuid.UUID `db:"passenger_id"`
	Type        string    `db:"type"`
	Number      string    `db:"number"`
}
