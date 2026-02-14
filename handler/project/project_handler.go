package project

import (
	"uplink-go/domain"
	"uplink-go/dto"
	"uplink-go/middleware"

	"github.com/gofiber/fiber/v3"
)

type CreateProjectRequest struct {
	Name string `json:"name"`
}

func (h *ProjectHandler) CreateProject(c fiber.Ctx) error {
	var req CreateProjectRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Name is required",
		})
	}

	userID, err := middleware.GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	project := &domain.Project{
		Name: req.Name,
		Users: []domain.User{
			{
				ID: userID,
			},
		},
	}

	if err := h.projectRepo.Create(project); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create project",
		})
	}

	return c.JSON(dto.ToProjectResponse(*project))
}

func (h *ProjectHandler) Projects(c fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	projects, err := h.userRepo.FindProjectsByUserID(userID)

	return c.JSON(dto.NewHydraResponse(
		dto.ToProjectsResponse(projects),
		1,
		10,
		len(projects),
	))
}
