package controller

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/1Storm3/avia-api/internal/converter"
	"github.com/1Storm3/avia-api/internal/dto"
	"github.com/1Storm3/avia-api/pkg/mistake"
)

type TicketController struct {
	ticketService TicketService
}

func NewTicketController(ticketService TicketService) *TicketController {
	return &TicketController{
		ticketService: ticketService,
	}
}

func (t *TicketController) GetAll(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	tickets, meta, err := t.ticketService.GetAll(c.Context(), dto.Pagination{
		Page:  &page,
		Limit: &limit,
	})

	if err != nil {
		return mistake.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"data": converter.DomainTicketsToResponse(tickets),
		"meta": meta,
	})
}

func (t *TicketController) Update(c *fiber.Ctx) error {

	var dtoTicket dto.UpdateTicketRequest
	if err := c.BodyParser(&dtoTicket); err != nil {
		return mistake.NewError(mistake.ErrInvalidRequestData, err.Error())
	}
	ticketID, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return mistake.NewError(mistake.ErrInvalidRequestData, err.Error())
	}
	dtoTicket.ID = ticketID

	domainTicket := converter.UpdateTicketRequestToDomain(dtoTicket)

	ticket, err := t.ticketService.Update(c.Context(), domainTicket)

	if err != nil {
		return mistake.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"data": converter.DomainTicketToResponse(ticket),
	})
}

func (t *TicketController) Delete(c *fiber.Ctx) error {
	ticketID := c.Params("id")

	id, err := uuid.Parse(ticketID)

	if err != nil {
		return mistake.NewError(mistake.ErrInvalidRequestData, err.Error())
	}
	err = t.ticketService.Delete(c.Context(), id)

	if err != nil {
		return mistake.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"data": "Билет успешно удален",
	})

}

func (t *TicketController) GetOne(c *fiber.Ctx) error {

	ticketID := c.Params("id")

	id, err := uuid.Parse(ticketID)

	if err != nil {
		return mistake.NewError(mistake.ErrInvalidRequestData, err.Error())
	}

	ticket, err := t.ticketService.GetOne(c.Context(), id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return mistake.NewError(mistake.ErrTicketNotFound, err.Error())
		}
		return mistake.HandleError(c, err)

	}

	return c.JSON(fiber.Map{
		"data": ticket,
	})
}

func (t *TicketController) GetPassengerReport(c *fiber.Ctx) error {
	passengerID := c.Query("passengerId")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	id, err := uuid.Parse(passengerID)

	if err != nil {
		return mistake.NewError(mistake.ErrInvalidRequestData, err.Error())
	}
	tickets, err := t.ticketService.GetPassengerReport(c.Context(), dto.GetPassengerReportRequest{
		PassengerID: id,
		StartDate:   startDate,
		EndDate:     endDate,
	})

	if err != nil {
		return mistake.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"data": tickets,
	})
}

func (t *TicketController) GetAllByPassenger(c *fiber.Ctx) error {
	passengerID := c.Params("id")

	id, err := uuid.Parse(passengerID)

	if err != nil {
		return mistake.NewError(mistake.ErrInvalidRequestData, err.Error())
	}

	tickets, err := t.ticketService.GetAllByPassenger(c.Context(), id)

	if err != nil {
		return mistake.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"data": converter.DomainTicketsToResponse(tickets),
	})
}
