package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/1Storm3/avia-api/internal/converter"
	"github.com/1Storm3/avia-api/internal/dto"
	"github.com/1Storm3/avia-api/internal/repo/model"
	"github.com/1Storm3/avia-api/pkg/gensqlc"
)

type TicketRepo struct {
	pool    *pgxpool.Pool
	queries *gensqlc.Queries
}

func NewTicketRepo(pool *pgxpool.Pool, queries *gensqlc.Queries) *TicketRepo {
	return &TicketRepo{
		pool:    pool,
		queries: queries,
	}
}

func (t *TicketRepo) GetAll(ctx context.Context, limit,
	offset int) ([]model.Ticket, int64, error) {

	rows, err := t.queries.GetAllTickets(ctx, gensqlc.GetAllTicketsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, 0, err
	}
	totalCount, err := t.queries.GetCountTickets(ctx)
	if err != nil {
		return nil, 0, err
	}
	return converter.SqlcTicketsToDomain(rows), totalCount, nil
}

func (t *TicketRepo) Update(ctx context.Context, ticket model.Ticket) (model.Ticket, error) {
	rows, err := t.queries.UpdateTicket(ctx, gensqlc.UpdateTicketParams{
		ID:              ticket.ID,
		OrderNumber:     ticket.OrderNumber,
		Departure:       ticket.Departure,
		Destination:     ticket.Destination,
		ServiceProvider: ticket.ServiceProvider,
		DepartureDate: converter.ParseTimeStringToTime(ticket.
			DepartureDate),
		ArrivalDate: converter.ParseTimeStringToTime(ticket.
			ArrivalDate),
		OrderDate: converter.ParseTimeStringToTime(ticket.
			OrderDate),
	})
	if err != nil {
		return model.Ticket{}, err
	}
	return converter.SqlcTicketToModel(rows), nil
}

func (t *TicketRepo) Delete(ctx context.Context, id uuid.UUID) error {
	err := t.queries.DeleteTicket(ctx, gensqlc.DeleteTicketParams{
		ID: id,
	})
	return err
}

func (t *TicketRepo) GetOne(ctx context.Context, id uuid.UUID) (dto.GetFullOneTicket, error) {
	rows, err := t.queries.GetOneTicket(ctx, gensqlc.GetOneTicketParams{
		ID: id,
	})

	if err != nil {
		return dto.GetFullOneTicket{}, err
	}

	return converter.SqlcGetOneRowToDto(rows), nil
}

func (t *TicketRepo) GetPassengerReport(ctx context.Context,
	dto dto.GetPassengerReportRequest) ([]dto.GetPassengerReportResponse, error) {
	result, err := t.queries.GetPassengerReport(ctx, gensqlc.GetPassengerReportParams{
		ID:            dto.PassengerID,
		DepartureDate: converter.ParseTimeStringToTime(dto.StartDate),
		OrderDate:     converter.ParseTimeStringToTime(dto.EndDate),
	})

	if err != nil {
		return nil, err
	}

	return converter.SqlcReportsToResponse(result), nil
}

func (t *TicketRepo) GetAllByPassenger(ctx context.Context, passengerID uuid.UUID) ([]model.Ticket, error) {
	result, err := t.queries.GetAllByPassenger(ctx, gensqlc.GetAllByPassengerParams{
		PassengerID: passengerID,
	})
	if err != nil {
		return nil, err
	}
	return converter.SqlcTicketsToDomain(result), nil
}
