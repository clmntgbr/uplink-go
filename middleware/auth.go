package middleware

import (
	"strings"

	"uplink-go/service"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type AuthMiddleware struct {
	authService *service.AuthService
}

func NewAuthMiddleware(authService *service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (m *AuthMiddleware) Protected() fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "unauthorized",
				"message": "Missing authorization header",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "unauthorized",
				"message": "Invalid authorization header format",
			})
		}

		if parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "unauthorized",
				"message": "Authorization scheme must be Bearer",
			})
		}

		tokenString := strings.TrimSpace(parts[1])
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "unauthorized",
				"message": "Token cannot be empty",
			})
		}

		claims, err := m.authService.ValidateToken(tokenString)
		if err != nil {
			statusCode := fiber.StatusUnauthorized
			message := "Invalid or expired token"
			switch {
			case strings.Contains(err.Error(), "expired"):
				message = "Token has expired"
			case strings.Contains(err.Error(), "malformed"):
				message = "Malformed token"
			case strings.Contains(err.Error(), "signature"):
				message = "Invalid token signature"
			}

			return c.Status(statusCode).JSON(fiber.Map{
				"error":   "unauthorized",
				"message": message,
			})
		}

		if claims == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "unauthorized",
				"message": "Invalid token claims",
			})
		}

		if claims.UserID == uuid.Nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "unauthorized",
				"message": "Missing user ID in token",
			})
		}

		c.Locals("user_id", claims.UserID)
		c.Locals("user_email", claims.Email)

		return c.Next()
	}
}

func GetUserID(c fiber.Ctx) (uuid.UUID, error) {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}
	return userID, nil
}

func GetUserEmail(c fiber.Ctx) (string, error) {
	email, ok := c.Locals("user_email").(string)
	if !ok {
		return "", fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}
	return email, nil
}