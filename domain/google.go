package domain

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

type GoogleResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Verified bool   `json:"verified_email"`
	Picture  string `json:"picture"`
}

type GoogleUsecase interface {
	GoogleConfig() *oauth2.Config
	GetToken(c *fiber.Ctx) (*oauth2.Token, error)
	GetProfile(token string) (*GoogleResponse, error)
}
