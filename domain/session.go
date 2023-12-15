package domain

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Session struct {
	Id        string    `json:"id" db:"id"`
	UserId    string    `json:"user_id" db:"user_id"`
	IpAddress string    `json:"ip_address" db:"ip_address"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ExpiredAt time.Time `json:"expired_at" db:"expired_at"`
}

type SessionRepository interface {
	Create(session *Session) error
	Get(id string) (*Session, error)
	Delete(id string) error
}

type SessionUsecase interface {
	Sign(id string) string
	Unsign(header string) (string, error)
	Create(userId string, ipAddress string) (*fiber.Cookie, error)
	Get(ssid string) (*Session, error)
	Delete(sid string) error
	Validate(ssid string) (bool, string, error)
}
