package converter

import (
	"github.com/1Storm3/avia-api/internal/domain"
	"github.com/1Storm3/avia-api/internal/dto"
	"github.com/1Storm3/avia-api/internal/repo/model"
	"github.com/1Storm3/avia-api/pkg/gensqlc"
)

func SqlcPassengerToModel(passenger *gensqlc.Passenger) model.Passenger {
	return model.Passenger{
		ID:         passenger.ID,
		FirstName:  passenger.FirstName,
		LastName:   passenger.LastName,
		MiddleName: passenger.MiddleName,
	}
}

func DomainPassengerToResponse(passenger domain.Passenger) dto.PassengerResponse {
	return dto.PassengerResponse{
		ID:         passenger.ID,
		FirstName:  passenger.FirstName,
		LastName:   passenger.LastName,
		MiddleName: passenger.MiddleName,
	}
}

func DomainPassengerToModel(passenger domain.Passenger) model.Passenger {
	return model.Passenger{
		ID:         passenger.ID,
		FirstName:  passenger.FirstName,
		LastName:   passenger.LastName,
		MiddleName: passenger.MiddleName,
	}
}

func ModelPassengerToDomain(passenger model.Passenger) domain.Passenger {
	return domain.Passenger{
		ID:         passenger.ID,
		FirstName:  passenger.FirstName,
		LastName:   passenger.LastName,
		MiddleName: passenger.MiddleName,
	}
}

func UpdatePassengerRequestToDomain(request dto.UpdatePassengerRequest) domain.Passenger {
	return domain.Passenger{
		ID:         request.ID,
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		MiddleName: request.MiddleName,
	}
}
