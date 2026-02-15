package endpoint

import (
	"uplink-go/ctxutil"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *EndpointHandler) Endpoints(c fiber.Ctx) error {
	ctx := c.Context()

	// Ajoutez string() autour des clés
	if userID := c.Locals(string(ctxutil.UserIDKey)); userID != nil {
		ctx = ctxutil.WithUserID(ctx, userID.(uuid.UUID))
	}

	if activeProjectID := c.Locals(string(ctxutil.ActiveProjectIDKey)); activeProjectID != nil {
		ctx = ctxutil.WithActiveProjectID(ctx, activeProjectID.(*uuid.UUID))
	}

	endpoints, err := h.endpointService.FindAll(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(endpoints)
}