package dto

import (
	"github.com/google/uuid"
)

type PassengerResponse struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	MiddleName string    `json:"middleName"`
}

type UpdatePassengerRequest struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"firstName" validate:"omitempty,min=2,max=100"`
	LastName   string    `json:"lastName" validate:"omitempty,min=2,max=100"`
	MiddleName string    `json:"middleName" validate:"omitempty,min=2,max=100"`
}
