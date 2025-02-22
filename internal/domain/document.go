package domain

import (
	"github.com/google/uuid"
)

type Document struct {
	ID          uuid.UUID
	PassengerID uuid.UUID
	Type        string
	Number      string
}
