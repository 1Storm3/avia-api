package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/1Storm3/avia-api/internal/converter"
	"github.com/1Storm3/avia-api/internal/domain"
)

type DocumentService struct {
	DocumentRepo DocumentRepo
}

func NewDocumentService(documentRepo DocumentRepo) *DocumentService {
	return &DocumentService{
		DocumentRepo: documentRepo,
	}
}

func (d *DocumentService) Delete(ctx context.Context, documentID uuid.UUID) error {
	return d.DocumentRepo.Delete(ctx, documentID)
}

func (d *DocumentService) Update(ctx context.Context,
	document domain.Document) (domain.Document, error) {

	documentModel := converter.DomainDocumentToModel(document)

	result, err := d.DocumentRepo.Update(ctx, documentModel)
	if err != nil {
		return domain.Document{}, err
	}

	return converter.ModelDocumentToDomain(result), nil
}

func (d *DocumentService) GetAllByPassengerID(ctx context.Context, passengerID uuid.UUID) ([]domain.Document, error) {

	result, err := d.DocumentRepo.GetAllByPassengerID(ctx, passengerID)

	if err != nil {
		return nil, err
	}

	return converter.ModelDocumentsToDomain(result), nil
}
