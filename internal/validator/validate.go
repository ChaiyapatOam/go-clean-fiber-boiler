package validator

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateBody(c *fiber.Ctx, Schema interface{}) error {
	c.BodyParser(Schema)
	validate := validator.New()
	err := validate.Struct(Schema)
	if err != nil {
		return errors.New("Invalid Body")
	}
	return nil
}
