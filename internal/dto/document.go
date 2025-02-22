package dto

import (
	"github.com/google/uuid"
)

type DocumentResponse struct {
	ID     uuid.UUID `json:"id"`
	Type   string    `json:"type"`
	Number string    `json:"number"`
}

type UpdateDocumentRequest struct {
	ID     uuid.UUID `json:"id"`
	Type   string    `json:"type" validate:"omitempty,min=2,max=100"`
	Number string    `json:"number" validate:"omitempty,min=2,max=100"`
}
