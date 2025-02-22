package converter

import (
	"github.com/1Storm3/avia-api/internal/domain"
	"github.com/1Storm3/avia-api/internal/dto"
	"github.com/1Storm3/avia-api/internal/repo/model"
	"github.com/1Storm3/avia-api/pkg/gensqlc"
)

func SqlcDocumentToModel(document *gensqlc.Document) model.Document {
	return model.Document{
		ID:          document.ID,
		PassengerID: *document.PassengerID,
		Type:        document.Type,
		Number:      document.Number,
	}
}

func SqlcDocumentsToModel(documents []*gensqlc.Document) []model.Document {
	var result []model.Document
	for _, document := range documents {
		result = append(result, SqlcDocumentToModel(document))
	}
	return result
}

func ModelDocumentsToDomain(documents []model.Document) []domain.Document {
	var result []domain.Document
	for _, document := range documents {
		result = append(result, ModelDocumentToDomain(document))
	}
	return result
}

func DomainDocumentToResponse(document domain.Document) dto.DocumentResponse {
	return dto.DocumentResponse{
		ID:     document.ID,
		Type:   document.Type,
		Number: document.Number,
	}
}

func DomainDocumentToModel(document domain.Document) model.Document {
	return model.Document{
		ID:          document.ID,
		PassengerID: document.PassengerID,
		Type:        document.Type,
		Number:      document.Number,
	}
}

func ModelDocumentToDomain(document model.Document) domain.Document {
	return domain.Document{
		ID:          document.ID,
		PassengerID: document.PassengerID,
		Type:        document.Type,
		Number:      document.Number,
	}
}

func DomainDocumentsToResponse(documents []domain.Document) []dto.DocumentResponse {
	var result []dto.DocumentResponse
	for _, document := range documents {
		result = append(result, DomainDocumentToResponse(document))
	}
	return result
}

func UpdateDocumentRequestToDomain(request dto.UpdateDocumentRequest) domain.Document {
	return domain.Document{
		ID:     request.ID,
		Type:   request.Type,
		Number: request.Number,
	}
}
