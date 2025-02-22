package service

import (
	"context"
	"math"

	"github.com/google/uuid"

	"github.com/1Storm3/avia-api/internal/converter"
	"github.com/1Storm3/avia-api/internal/domain"
	"github.com/1Storm3/avia-api/internal/dto"
)

type TicketService struct {
	ticketRepo TicketRepo
}

func NewTicketService(ticketRepo TicketRepo) *TicketService {
	return &TicketService{
		ticketRepo: ticketRepo,
	}
}

func (t *TicketService) GetAll(ctx context.Context, pagination dto.Pagination) (
	[]domain.Ticket, dto.PaginationMeta, error) {

	offset := (*pagination.Page - 1) * *pagination.Limit

	tickets, count, err := t.ticketRepo.GetAll(ctx, *pagination.Limit,
		offset)

	if err != nil {
		return nil, dto.PaginationMeta{}, err
	}

	return converter.ModelTicketsToDomain(tickets), dto.PaginationMeta{
		Limit: *pagination.Limit,
		Page:  *pagination.Page,
		Total: int(count),
		TotalPages: int(math.Ceil(float64(count) / float64(
			*pagination.Limit))),
	}, nil
}

func (t *TicketService) Update(ctx context.Context, ticket domain.Ticket) (domain.Ticket, error) {

	ticketModel := converter.DomainTicketToModel(ticket)

	result, err := t.ticketRepo.Update(ctx, ticketModel)

	if err != nil {
		return domain.Ticket{}, err
	}

	return converter.ModelTicketToDomain(result), nil
}

func (t *TicketService) Delete(ctx context.Context, ticketID uuid.UUID) error {
	return t.ticketRepo.Delete(ctx, ticketID)
}

func (t *TicketService) GetOne(ctx context.Context, ticketID uuid.UUID) (dto.GetFullOneTicket,
	error) {
	result, err := t.ticketRepo.GetOne(ctx, ticketID)

	if err != nil {
		return dto.GetFullOneTicket{}, err
	}

	return result, nil
}

func (t *TicketService) GetPassengerReport(ctx context.Context,
	dto dto.GetPassengerReportRequest) ([]dto.GetPassengerReportResponse, error) {
	return t.ticketRepo.GetPassengerReport(ctx, dto)
}

func (t *TicketService) GetAllByPassenger(ctx context.Context, passengerID uuid.UUID) ([]domain.Ticket, error) {
	result, err := t.ticketRepo.GetAllByPassenger(ctx, passengerID)

	if err != nil {
		return nil, err
	}

	return converter.ModelTicketsToDomain(result), nil
}
