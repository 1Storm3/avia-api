package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/1Storm3/avia-api/internal/converter"
	"github.com/1Storm3/avia-api/internal/repo/model"
	"github.com/1Storm3/avia-api/pkg/gensqlc"
)

type DocumentRepo struct {
	pool    *pgxpool.Pool
	queries *gensqlc.Queries
}

func NewDocumentRepo(pool *pgxpool.Pool, queries *gensqlc.Queries) *DocumentRepo {
	return &DocumentRepo{
		pool:    pool,
		queries: queries,
	}
}

func (d *DocumentRepo) Delete(ctx context.Context, id uuid.UUID) error {
	err := d.queries.DeleteDocument(ctx, gensqlc.DeleteDocumentParams{
		ID: id,
	})
	return err
}

func (d *DocumentRepo) Update(ctx context.Context, document model.Document) (model.Document, error) {
	rows, err := d.queries.UpdateDocument(ctx, gensqlc.UpdateDocumentParams{
		ID:     document.ID,
		Type:   document.Type,
		Number: document.Number,
	})

	if err != nil {
		return model.Document{}, err
	}
	return converter.SqlcDocumentToModel(rows), nil
}

func (d *DocumentRepo) GetAllByPassengerID(ctx context.Context, passengerID uuid.UUID) ([]model.Document, error) {
	rows, err := d.queries.GetDocumentsByPassenger(ctx,
		gensqlc.GetDocumentsByPassengerParams{
			PassengerID: &passengerID,
		})
	if err != nil {
		return nil, err
	}

	return converter.SqlcDocumentsToModel(rows), nil
}
