package usecase

import (
	"time"

	"github.com/chaiyapatoam/go-clean-fiber-boiler/domain"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/payload"
	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (u *UserUsecase) Create(body payload.CreateUser) error {
	id := uuid.NewString()

	user := &domain.User{
		Id:         id,
		Email:      body.Email,
		FirstName:  body.FirstName,
		LastName:   body.LastName,
		ProfileUrl: body.ProfileUrl,
		CreatedAt:  time.Now(),
	}
	
	err := u.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) Get(id string) (*domain.User, error) {
	user, err := u.userRepository.Get(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
