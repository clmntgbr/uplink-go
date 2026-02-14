package user

import (
	"uplink-go/dto"
	"uplink-go/middleware"

	"github.com/gofiber/fiber/v3"
)

func (h *UserHandler) User(c fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	user, err := h.userRepo.FindByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(dto.ToUserResponse(user))
}
