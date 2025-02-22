package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/1Storm3/avia-api/internal/converter"
	"github.com/1Storm3/avia-api/internal/domain"
)

type PassengerService struct {
	PassengerRepo PassengerRepo
}

func NewPassengerService(passengerRepo PassengerRepo) *PassengerService {
	return &PassengerService{
		PassengerRepo: passengerRepo,
	}
}

func (p *PassengerService) Delete(ctx context.Context,
	passengerID uuid.UUID) error {
	return p.PassengerRepo.Delete(ctx, passengerID)
}

func (p *PassengerService) Update(ctx context.Context,
	passenger domain.Passenger) (domain.Passenger, error) {
	passengerModel := converter.DomainPassengerToModel(passenger)

	result, err := p.PassengerRepo.Update(ctx, passengerModel)

	if err != nil {
		return domain.Passenger{}, err
	}

	return converter.ModelPassengerToDomain(result), nil
}
