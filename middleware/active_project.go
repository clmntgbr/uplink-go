package middleware

import (
	"github.com/gofiber/fiber/v3"
	"uplink-go/ctxutil"
	"uplink-go/repository"
)

func InjectActiveProject(projectRepo *repository.ProjectRepository) fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, ok := ctxutil.GetUserIDFromContext(c.Context())
		if !ok {
			return c.Next()
		}
		activeProjectID, err := projectRepo.FindActiveProject(c.Context(), userID)
		if err == nil && activeProjectID != nil {
			c.Locals(string(ctxutil.ActiveProjectIDKey), activeProjectID)
		}

		return c.Next()
	}
}