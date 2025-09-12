package approval

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomamac/tbtest/backend/internal/entity"
)

type HttpHandler struct {
	service IService
}

func RegisterHandlers(g fiber.Router, service IService) {
	handler := &HttpHandler{service}

	g.Get("/", handler.get)
	g.Put("/", handler.update)
}

type UpdateApprovalRequest struct {
	IDs         []uint                `json:"ids"`
	Description string                `json:"description"`
	Status      entity.ApprovalStatus `json:"status"`
}

func (h *HttpHandler) get(c *fiber.Ctx) error {
	approvals, err := h.service.Get()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Not found"})
	}

	return c.JSON(approvals)
}

func (h *HttpHandler) update(c *fiber.Ctx) error {
	var approvals UpdateApprovalRequest

	if err := c.BodyParser(&approvals); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Wrong format"})
	}

	if err := h.service.Update(approvals); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Success"})
}
