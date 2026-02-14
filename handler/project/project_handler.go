package project

import (
	"uplink-go/middleware"
	"uplink-go/service"

	"github.com/gofiber/fiber/v3"
)

func (h *ProjectHandler) CreateProject(c fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var req service.CreateProjectRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	resp, err := h.createProjectService.Create(userID, req)
	if err != nil {
		if err.Error() == "name is required" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create project",
		})
	}

	return c.JSON(resp)
}

func (h *ProjectHandler) Projects(c fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	resp, err := h.getProjectsService.Projects(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch projects",
		})
	}

	return c.JSON(resp)
}
