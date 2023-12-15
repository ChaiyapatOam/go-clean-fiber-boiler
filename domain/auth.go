package domain

import (
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/payload"
	"github.com/gofiber/fiber/v2"
)

type AuthUsecase interface {
	GoogleLogin(c *fiber.Ctx) (*User, *fiber.Cookie, error)
	Login(body payload.Login, ipAddress string) (*User, *fiber.Cookie, error)
	Logout(sid string) (*fiber.Cookie, error)
}
