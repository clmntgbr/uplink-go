package project

import (
	"errors"
	"uplink-go/dto"
	apperrors "uplink-go/errors"
	"uplink-go/middleware"
	"uplink-go/validator"

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

	var input dto.CreateInput

	if err := c.Bind().Body(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}

	if err := validator.ValidateStruct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  validator.FormatValidationErrors(err),
		})
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

	project, err := h.projectService.FindById(c.Context(), userID, projectUUID)
	if err != nil {
		if errors.Is(err, apperrors.ErrProjectNotFound) {
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
