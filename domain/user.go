package domain

import (
	"time"

	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/payload"
)

type User struct {
	Id         string    `json:"id" db:"id"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"-" db:"password"`
	FirstName  string    `json:"firstName" db:"first_name"`
	LastName   string    `json:"lastName" db:"last_name"`
	Phone      string    `json:"phone" db:"phone"`
	Provider   string    `json:"provider" db:"provider"`
	ProfileUrl string    `json:"profileUrl" db:"profile_url"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
}

type UserRepository interface {
	Create(user *User) error
	Get(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(user *User) error
	ChangePassword(userId string, password string) error
}

type UserUsecase interface {
	Create(data payload.Register) error
	CreateFromGoogle(profile *GoogleResponse) (*User, error)
	Get(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(id string, data payload.UpdateUser) error
	ChangePassword(userId string, data payload.ChangePassword) error
}
