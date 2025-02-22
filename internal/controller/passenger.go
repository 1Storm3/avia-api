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

type PassengerController struct {
	passengerService PassengerService
}

func NewPassengerController(passengerService PassengerService) *PassengerController {
	return &PassengerController{
		passengerService: passengerService,
	}
}

func (p *PassengerController) Delete(c *fiber.Ctx) error {
	passengerID := c.Params("id")

	id, err := uuid.Parse(passengerID)

	if err != nil {
		return mistake.NewError(mistake.ErrInvalidRequestData, err.Error())
	}

	err = p.passengerService.Delete(c.Context(), id)

	if err != nil {
		return mistake.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"data": "Пассажир успешно удален",
	})
}

func (p *PassengerController) Update(c *fiber.Ctx) error {
	var dtoPassenger dto.UpdatePassengerRequest
	if err := c.BodyParser(&dtoPassenger); err != nil {
		return mistake.NewError(mistake.ErrInvalidRequestData, err.Error())
	}

	passengerID, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return mistake.NewError(mistake.ErrInvalidRequestData, err.Error())
	}
	dtoPassenger.ID = passengerID

	passengerDomain := converter.UpdatePassengerRequestToDomain(
		dtoPassenger)

	passenger, err := p.passengerService.Update(c.Context(),
		passengerDomain)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return mistake.NewError(mistake.ErrPassengerNotFound, err.Error())
		}
		return mistake.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"data": converter.DomainPassengerToResponse(passenger),
	})
}
