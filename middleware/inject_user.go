package middleware

import (
	"uplink-go/ctxutil"

	"github.com/gofiber/fiber/v3"
)

func InjectUserContext() fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := GetUserID(c)
		if err != nil {
			return c.Next()
		}

		c.Locals(ctxutil.UserIDKey, userID)

		return c.Next()
	}
}
