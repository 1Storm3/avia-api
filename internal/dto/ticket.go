package dto

import (
	"github.com/google/uuid"
)

type UpdateTicketRequest struct {
	ID              uuid.UUID `json:"id"`
	OrderNumber     string    `json:"orderNumber" validate:"omitempty,min=1,max=50"`
	Departure       string    `json:"departure" validate:"omitempty,min=2,max=100"`
	Destination     string    `json:"destination" validate:"omitempty,min=2,max=100"`
	ServiceProvider string    `json:"serviceProvider" validate:"omitempty,min=2,max=100"`
	DepartureDate   string    `json:"departureDate" validate:"omitempty,datetime=2006-01-02T15:04:05"`
	ArrivalDate     string    `json:"arrivalDate" validate:"omitempty,datetime=2006-01-02T15:04:05"`
	OrderDate       string    `json:"orderDate" validate:"omitempty,datetime=2006-01-02T15:04:05"`
}

type UpdateTicketResponse struct {
	ID              uuid.UUID `json:"id"`
	OrderNumber     string    `json:"orderNumber"`
	Departure       string    `json:"departure"`
	Destination     string    `json:"destination"`
	ServiceProvider string    `json:"serviceProvider"`
	DepartureDate   string    `json:"departureDate"`
	ArrivalDate     string    `json:"arrivalDate"`
	OrderDate       string    `json:"orderDate"`
}

type TicketResponse struct {
	ID              uuid.UUID `json:"id"`
	OrderNumber     string    `json:"orderNumber"`
	Departure       string    `json:"departure"`
	Destination     string    `json:"destination"`
	ServiceProvider string    `json:"serviceProvider"`
	DepartureDate   string    `json:"departureDate"`
	ArrivalDate     string    `json:"arrivalDate"`
	OrderDate       string    `json:"orderDate"`
}

type GetFullOneTicket struct {
	Ticket    TicketResponse    `json:"ticket"`
	Passenger PassengerResponse `json:"passenger"`
	Document  DocumentResponse  `json:"document"`
}

type GetPassengerReportRequest struct {
	PassengerID uuid.UUID `json:"passengerId" validate:"required"`
	StartDate   string    `json:"startDate" validate:"required"`
	EndDate     string    `json:"endDate" validate:"required"`
}

type GetPassengerReportResponse struct {
	OrderDate       string `json:"orderDate"`
	DepartureDate   string `json:"departureDate"`
	OrderNumber     string `json:"orderNumber"`
	Departure       string `json:"departure"`
	Destination     string `json:"destination"`
	ServiceProvided bool   `json:"serviceProvided"`
}
