package dto

import (
	"github.com/google/uuid"
)

type PassengerResponse struct {
	ID         uuid.UUID          `json:"id"`
	FirstName  string             `json:"firstName"`
	LastName   string             `json:"lastName"`
	MiddleName string             `json:"middleName"`
	Documents  []DocumentResponse `json:"documents"`
}

type UpdatePassengerRequest struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"firstName" validate:"required,min=2,max=100"`
	LastName   string    `json:"lastName" validate:"required,min=2,max=100"`
	MiddleName string    `json:"middleName" validate:"required,min=2,max=100"`
}
