package controller

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/1Storm3/avia-api/internal/converter"
	"github.com/1Storm3/avia-api/internal/dto"
	"github.com/1Storm3/avia-api/pkg/mistake"
)

type DocumentController struct {
	DocumentService DocumentService
}

func NewDocumentController(documentService DocumentService) *DocumentController {
	return &DocumentController{
		DocumentService: documentService,
	}
}

func (d *DocumentController) Delete(c *fiber.Ctx) error {
	documentID := c.Params("id")

	id, err := uuid.Parse(documentID)

	if err != nil {
		return mistake.NewError(mistake.ErrInvalidRequestData, err.Error())
	}

	err = d.DocumentService.Delete(c.Context(), id)

	if err != nil {
		return mistake.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"data": "Документ успешно удален",
	})
}

func (d *DocumentController) Update(c *fiber.Ctx) error {
	var dtoDocument dto.UpdateDocumentRequest

	if err := c.BodyParser(&dtoDocument); err != nil {
		return mistake.NewError(mistake.ErrInvalidRequestData, err.Error())
	}

	documentID, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return mistake.NewError(mistake.ErrInvalidRequestData, err.Error())
	}
	dtoDocument.ID = documentID

	documentDomain := converter.UpdateDocumentRequestToDomain(dtoDocument)

	document, err := d.DocumentService.Update(c.Context(), documentDomain)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return mistake.NewError(mistake.ErrDocumentNotFound, err.Error())
		}
		return mistake.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"data": converter.DomainDocumentToResponse(document),
	})
}

func (d *DocumentController) GetAllByPassengerID(c *fiber.Ctx) error {
	passengerID := c.Params("passengerId")

	id, err := uuid.Parse(passengerID)

	if err != nil {
		return mistake.NewError(mistake.ErrInvalidRequestData, err.Error())
	}

	documents, err := d.DocumentService.GetAllByPassengerID(c.Context(), id)

	if err != nil {
		return mistake.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"data": converter.DomainDocumentsToResponse(documents),
	})
}
