package http

import (
	"github.com/gofiber/fiber/v2"
)

type TicketController interface {
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	GetPassengerReport(c *fiber.Ctx) error
	GetAllByPassenger(c *fiber.Ctx) error
}

type PassengerController interface {
	Delete(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}

type DocumentController interface {
	Delete(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	GetAllByPassengerID(c *fiber.Ctx) error
}
