package domain

import (
	"time"

	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/payload"
)

type User struct {
	Id         string    `json:"id" db:"id"`
	Email      string    `json:"email" db:"email"`
	FirstName  string    `json:"firstName" db:"first_name"`
	LastName   string    `json:"lastName" db:"last_name"`
	ProfileUrl string    `json:"profileUrl" db:"profile_url"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
}

type UserRepository interface {
	Create(user *User) error
	Get(email string) (*User, error)
}

type UserUsecase interface {
	Create(data payload.CreateUser) error
	Get(email string) (*User, error)
}
