package project

import (
	"uplink-go/dto"
	"uplink-go/middleware"

	"github.com/gofiber/fiber/v3"
)

func (h *ProjectHandler) Projects(c fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	projects, err := h.userRepo.FindProjectsByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Projects not found",
		})
	}

	return c.JSON(dto.NewHydraResponse(
		dto.ToProjectsResponse(projects),
		1,
		10,
		len(projects),
	))
}
