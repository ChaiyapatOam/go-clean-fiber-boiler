package middleware

import (
	"github.com/chaiyapatoam/go-clean-fiber-boiler/domain"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/response"
	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleWare(sessionUsecase domain.SessionUsecase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ssid := c.Cookies("ssid")
		if ssid == "" {
			return response.ErrorResponse(c, 403, "Missing Cookie")
		}

		check, userId, err := sessionUsecase.Validate(ssid)

		if !check || err != nil {
			return response.ErrorResponse(c, 401, "Unauthorized")
		}

		c.Locals("userId", userId)
		return c.Next()
	}
}

func GetUser(c *fiber.Ctx) string {
	userId := c.Locals("userId").(string)
	return userId
}
