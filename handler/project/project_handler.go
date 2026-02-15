package project

import (
	"errors"
	"uplink-go/dto"
	apperrors "uplink-go/errors"
	"uplink-go/middleware"
	"uplink-go/validator"
	"uplink-go/ctxutil"

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
	ctx := c.Context()

	if userID := c.Locals(string(ctxutil.UserIDKey)); userID != nil {
		ctx = ctxutil.WithUserID(ctx, userID.(uuid.UUID))
	}

	if activeProjectID := c.Locals(string(ctxutil.ActiveProjectIDKey)); activeProjectID != nil {
		ctx = ctxutil.WithActiveProjectID(ctx, activeProjectID.(*uuid.UUID))
	}

	projects, err := h.projectService.FindAll(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(projects)
}

func (h *ProjectHandler) ProjectById(c fiber.Ctx) error {
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
	ctx := c.Context()

	if userID := c.Locals(string(ctxutil.UserIDKey)); userID != nil {
		ctx = ctxutil.WithUserID(ctx, userID.(uuid.UUID))
	}
	if activeProjectID := c.Locals(string(ctxutil.ActiveProjectIDKey)); activeProjectID != nil {
		ctx = ctxutil.WithActiveProjectID(ctx, activeProjectID.(*uuid.UUID))
	}

	project, err := h.projectService.FindById(ctx, projectUUID)
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
