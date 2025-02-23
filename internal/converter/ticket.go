package converter

import (
	"github.com/1Storm3/avia-api/internal/domain"
	"github.com/1Storm3/avia-api/internal/dto"
	"github.com/1Storm3/avia-api/internal/repo/model"
	"github.com/1Storm3/avia-api/pkg/gensqlc"
)

func SqlcTicketToModel(ticket *gensqlc.Ticket) model.Ticket {
	return model.Ticket{
		ID:              ticket.ID,
		OrderNumber:     ticket.OrderNumber,
		Departure:       ticket.Departure,
		Destination:     ticket.Destination,
		ServiceProvider: ticket.ServiceProvider,
		DepartureDate:   ticket.DepartureDate.String(),
		ArrivalDate:     ticket.ArrivalDate.String(),
		OrderDate:       ticket.OrderDate.String(),
	}
}

func SqlcTicketsToDomain(tickets []*gensqlc.Ticket) []model.Ticket {
	var result []model.Ticket
	for _, ticket := range tickets {
		result = append(result, SqlcTicketToModel(ticket))
	}
	return result
}

func DomainTicketToModel(ticket domain.Ticket) model.Ticket {
	return model.Ticket{
		ID:              ticket.ID,
		OrderNumber:     ticket.OrderNumber,
		Departure:       ticket.Departure,
		Destination:     ticket.Destination,
		ServiceProvider: ticket.ServiceProvider,
		DepartureDate:   ticket.DepartureDate,
		ArrivalDate:     ticket.ArrivalDate,
		OrderDate:       ticket.OrderDate,
	}
}

func ModelTicketToDomain(ticket model.Ticket) domain.Ticket {
	return domain.Ticket{
		ID:              ticket.ID,
		OrderNumber:     ticket.OrderNumber,
		Departure:       ticket.Departure,
		Destination:     ticket.Destination,
		ServiceProvider: ticket.ServiceProvider,
		DepartureDate:   ticket.DepartureDate,
		ArrivalDate:     ticket.ArrivalDate,
		OrderDate:       ticket.OrderDate,
	}
}

func ModelTicketsToDomain(tickets []model.Ticket) []domain.Ticket {
	var result []domain.Ticket
	for _, ticket := range tickets {
		result = append(result, ModelTicketToDomain(ticket))
	}
	return result
}

func UpdateTicketRequestToDomain(request dto.UpdateTicketRequest) domain.Ticket {
	return domain.Ticket{
		ID:              request.ID,
		OrderNumber:     request.OrderNumber,
		Departure:       request.Departure,
		Destination:     request.Destination,
		ServiceProvider: request.ServiceProvider,
		DepartureDate:   request.DepartureDate,
		ArrivalDate:     request.ArrivalDate,
		OrderDate:       request.OrderDate,
	}
}

func DomainTicketsToResponse(tickets []domain.Ticket) []dto.TicketResponse {
	var result []dto.TicketResponse
	for _, ticket := range tickets {
		result = append(result, DomainTicketToResponse(ticket))
	}
	return result
}

func DomainTicketToResponse(ticket domain.Ticket) dto.TicketResponse {
	return dto.TicketResponse{
		ID:              ticket.ID,
		OrderNumber:     ticket.OrderNumber,
		Departure:       ticket.Departure,
		Destination:     ticket.Destination,
		ServiceProvider: ticket.ServiceProvider,
		DepartureDate:   ticket.DepartureDate,
		ArrivalDate:     ticket.ArrivalDate,
		OrderDate:       ticket.OrderDate,
	}
}

func SqlcReportToResponse(report *gensqlc.GetPassengerReportRow) dto.GetPassengerReportResponse {
	return dto.GetPassengerReportResponse{
		OrderDate:       report.OrderDate.String(),
		DepartureDate:   report.DepartureDate.String(),
		OrderNumber:     report.OrderNumber,
		Departure:       report.Departure,
		Destination:     report.Destination,
		ServiceProvided: report.ServiceProvided,
	}
}

func SqlcReportsToResponse(reports []*gensqlc.GetPassengerReportRow) []dto.GetPassengerReportResponse {
	var result []dto.GetPassengerReportResponse
	for _, report := range reports {
		result = append(result, SqlcReportToResponse(report))
	}
	return result
}
