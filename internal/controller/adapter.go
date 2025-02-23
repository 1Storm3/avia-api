package controller

import (
	"context"

	"github.com/google/uuid"

	"github.com/1Storm3/avia-api/internal/domain"
	"github.com/1Storm3/avia-api/internal/dto"
)

type TicketService interface {
	GetAll(ctx context.Context, pagination dto.Pagination) (
		[]domain.Ticket, dto.PaginationMeta, error)
	Update(ctx context.Context, ticket domain.Ticket) (domain.Ticket, error)
	Delete(ctx context.Context, ticketID uuid.UUID) error
	GetOne(ctx context.Context, ticketID uuid.UUID) (*dto.GetFullOneTicket, error)
	GetPassengerReport(ctx context.Context,
		dto dto.GetPassengerReportRequest) ([]dto.GetPassengerReportResponse, error)
	GetAllByPassenger(ctx context.Context, passengerID uuid.UUID) ([]domain.Ticket, error)
}

type PassengerService interface {
	Delete(ctx context.Context, passengerID uuid.UUID) error
	Update(ctx context.Context, passenger domain.Passenger) (domain.Passenger, error)
}

type DocumentService interface {
	Delete(ctx context.Context, documentID uuid.UUID) error
	Update(ctx context.Context, document domain.Document) (domain.Document, error)
	GetAllByPassengerID(ctx context.Context, passengerID uuid.UUID) ([]domain.Document, error)
}
