package project

import (
	"uplink-go/service/project"

	"github.com/gofiber/fiber/v3"
)

func (h *ProjectHandler) CreateProject(c fiber.Ctx) error {
    var input project.CreateInput

    if err := c.Bind().Body(&input); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "invalid body")
    }

    projectCreated, err := h.projectService.Create(
        c.Context(),
        input,
    )
    if err != nil {
        return err
    }

    return c.Status(fiber.StatusCreated).JSON(projectCreated)
}


func (h *ProjectHandler) Projects(c fiber.Ctx) error {
    projects, err := h.projectService.FindAll(
        c.Context(),
    )
    if err != nil {
        return err
    }

    return c.JSON(projects)
}


