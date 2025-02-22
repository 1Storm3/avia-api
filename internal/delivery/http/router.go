package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/1Storm3/avia-api/internal/dto"
	"github.com/1Storm3/avia-api/pkg/middleware"
)

type Router struct {
	ticketController    TicketController
	passengerController PassengerController
	documentController  DocumentController
}

func NewRouter(ticketController TicketController,
	passengerController PassengerController,
	documentController DocumentController) *Router {
	return &Router{
		ticketController:    ticketController,
		passengerController: passengerController,
		documentController:  documentController,
	}
}

func (r *Router) RegisterRoutes(app fiber.Router) {
	apiRoute := app.Group("/api")

	// Билеты
	ticketRoute := apiRoute.Group("/ticket")

	ticketRoute.Get("/", middleware.
		ValidateMiddleware[dto.Pagination](middleware.FromQuery),
		r.ticketController.GetAll)

	ticketRoute.Patch("/:id", middleware.ValidateMiddleware[dto.
		UpdateTicketRequest](
		middleware.FromBody), r.ticketController.Update)

	ticketRoute.Get("/report", middleware.ValidateMiddleware[dto.GetPassengerReportRequest](
		middleware.
			FromQuery),
		r.ticketController.GetPassengerReport)

	ticketRoute.Get("/passenger/:id", r.ticketController.GetAllByPassenger)

	ticketRoute.Get("/:id", r.ticketController.GetOne)

	ticketRoute.Delete("/:id", r.ticketController.Delete)

	// Пассажиры
	passengerRoute := apiRoute.Group("/passenger")

	passengerRoute.Delete("/:id", r.passengerController.Delete)

	passengerRoute.Patch("/:id", middleware.ValidateMiddleware[dto.
		UpdatePassengerRequest](
		middleware.FromBody), r.passengerController.Update)

	// Документы
	documentRoute := apiRoute.Group("/document")

	documentRoute.Delete("/:id", r.documentController.Delete)

	documentRoute.Get("/passenger/:passengerId",
		r.documentController.GetAllByPassengerID)

	documentRoute.Patch("/:id", middleware.ValidateMiddleware[dto.
		UpdateDocumentRequest](
		middleware.FromBody), r.documentController.Update)
}
