package response

import "github.com/gofiber/fiber/v2"

func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{"success": false, "message": message})
}

func SuccessResponse(c *fiber.Ctx, message string) error {
	return c.Status(200).JSON(fiber.Map{"success": true, "message": message})
}

func SuccessDataResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(200).JSON(fiber.Map{"success": true, "data": data})
}
