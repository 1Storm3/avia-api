package repo

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/1Storm3/avia-api/internal/converter"
	"github.com/1Storm3/avia-api/internal/repo/model"
	"github.com/1Storm3/avia-api/pkg/gensqlc"
)

type PassengerRepo struct {
	pool    *pgxpool.Pool
	queries *gensqlc.Queries
}

func NewPassengerRepo(pool *pgxpool.Pool, queries *gensqlc.Queries) *PassengerRepo {
	return &PassengerRepo{
		pool:    pool,
		queries: queries,
	}
}

func (p *PassengerRepo) Delete(ctx context.Context, passengerID uuid.UUID) error {
	err := p.queries.DeletePassenger(ctx, gensqlc.DeletePassengerParams{
		ID: passengerID,
	})
	log.Println(err)
	return err
}

func (p *PassengerRepo) Update(ctx context.Context, passenger model.Passenger) (model.Passenger, error) {
	rows, err := p.queries.UpdatePassenger(ctx, gensqlc.UpdatePassengerParams{
		ID:         passenger.ID,
		FirstName:  passenger.FirstName,
		LastName:   passenger.LastName,
		MiddleName: passenger.MiddleName,
	})
	if err != nil {
		return model.Passenger{}, err
	}
	return converter.SqlcPassengerToModel(rows), nil
}
