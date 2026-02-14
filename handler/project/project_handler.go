package project

import (
	"uplink-go/middleware"
	"uplink-go/service/project"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *ProjectHandler) CreateProject(c fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var input project.CreateInput

	if err := c.Bind().Body(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}

	projectCreated, err := h.projectService.Create(
		c.Context(),
		input,
		userID,
	)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(projectCreated)
}

func (h *ProjectHandler) Projects(c fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	projects, err := h.projectService.FindAll(
		c.Context(),
		userID,
	)
	if err != nil {
		return err
	}

	return c.JSON(projects)
}

func (h *ProjectHandler) ProjectById(c fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	projectID := c.Params("id")
	if projectID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid project ID",
		})
	}

	projectUUID, err := uuid.Parse(projectID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid project ID format",
		})
	}

	project, err := h.projectService.FindById(
		c.Context(),
		userID,
		projectUUID,
	)
	if err != nil {
		if err.Error() == "project not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Project not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(project)
}
