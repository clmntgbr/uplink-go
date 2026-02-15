package middleware

import (
	"uplink-go/ctxutil"
	"uplink-go/repository"

	"github.com/gofiber/fiber/v3"
)

func InjectActiveProject(projectRepo *repository.ProjectRepository) fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := GetUserID(c)
		if err != nil {
			return c.Next()
		}

		activeProjectID, err := projectRepo.FindActiveProject(c.Context(), userID)
		if err == nil && activeProjectID != nil {
			c.Locals(ctxutil.ActiveProjectIDKey, activeProjectID)
		}

		return c.Next()
	}
}
