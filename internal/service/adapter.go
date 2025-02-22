package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/1Storm3/avia-api/internal/dto"
	"github.com/1Storm3/avia-api/internal/repo/model"
)

type TicketRepo interface {
	GetAll(ctx context.Context, limit, offset int) ([]model.Ticket, int64, error)
	Update(ctx context.Context, ticket model.Ticket) (model.Ticket, error)
	Delete(ctx context.Context, ticketID uuid.UUID) error
	GetOne(ctx context.Context, ticketID uuid.UUID) (dto.GetFullOneTicket, error)
	GetPassengerReport(ctx context.Context,
		dto dto.GetPassengerReportRequest) ([]dto.GetPassengerReportResponse, error)
	GetAllByPassenger(ctx context.Context, passengerID uuid.UUID) ([]model.Ticket, error)
}

type PassengerRepo interface {
	Delete(ctx context.Context, passengerID uuid.UUID) error
	Update(ctx context.Context, passenger model.Passenger) (model.Passenger, error)
}

type DocumentRepo interface {
	Delete(ctx context.Context, documentID uuid.UUID) error
	Update(ctx context.Context, document model.Document) (model.Document, error)
	GetAllByPassengerID(ctx context.Context, passengerID uuid.UUID) ([]model.Document, error)
}
