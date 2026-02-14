package user

import (
	"uplink-go/middleware"

	"github.com/gofiber/fiber/v3"
)

func (h *UserHandler) User(c fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	user, err := h.userRepo.FindByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"id": user.ID,
		"email": user.Email,
	})
}
