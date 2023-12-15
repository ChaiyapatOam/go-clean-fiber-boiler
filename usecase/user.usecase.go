package usecase

import (
	"errors"
	"time"

	"github.com/chaiyapatoam/go-clean-fiber-boiler/domain"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/payload"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepository domain.UserRepository
	sessionUsecase domain.SessionUsecase
}

func NewUserUsecase(userRepository domain.UserRepository, sessionUsecase domain.SessionUsecase) domain.UserUsecase {
	return &UserUsecase{userRepository: userRepository, sessionUsecase: sessionUsecase}
}

func (u *UserUsecase) Create(body payload.Register) error {
	id := uuid.NewString()

	user, err := u.FindByEmail(body.Email)
	if err != nil {
		return err
	} else if user != nil {
		return errors.New("user already exist")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return errors.New("cannot create user with invalid password")
	}

	user = &domain.User{
		Id:         id,
		Email:      body.Email,
		Password:   string(passwordHash),
		FirstName:  body.FirstName,
		LastName:   body.LastName,
		Provider:   "",
		ProfileUrl: body.ProfileUrl,
		CreatedAt:  time.Now(),
	}

	err = u.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) CreateFromGoogle(profile *domain.GoogleResponse) (*domain.User, error) {
	id := uuid.NewString()

	user := &domain.User{
		Id:         id,
		Email:      profile.Email,
		Password:   "",
		FirstName:  profile.Name,
		LastName:   "",
		Provider:   "GOOGLE",
		ProfileUrl: profile.Picture,
		CreatedAt:  time.Now(),
	}

	err := u.userRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) Get(id string) (*domain.User, error) {
	user, err := u.userRepository.Get(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecase) FindByEmail(email string) (*domain.User, error) {
	user, err := u.userRepository.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecase) Update(id string, data payload.UpdateUser) error {
	user := &domain.User{
		Id:         id,
		FirstName:  data.FirstName,
		LastName:   data.LastName,
		Phone:      data.Phone,
		ProfileUrl: data.ProfileUrl,
	}

	err := u.userRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) ChangePassword(id string, data payload.ChangePassword) error {
	user, err := u.userRepository.Get(id)
	if err != nil {
		return errors.New("cant update user password")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.OldPassword)); err != nil {
		return errors.New("old password incorrect")
	}

	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(data.NewPassword), 10)
	if err != nil {
		return errors.New("cannot create user with invalid password")
	}

	err = u.userRepository.ChangePassword(user.Id, string(newPasswordHash))
	if err != nil {
		return err
	}

	// TODO: Delete session by userId
	// err = u.sessionUsecase.Delete()
	return nil
}
