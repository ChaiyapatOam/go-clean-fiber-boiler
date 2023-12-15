package usecase

import (
	"errors"
	"time"

	"github.com/chaiyapatoam/go-clean-fiber-boiler/domain"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/payload"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	googleUsecase  domain.GoogleUsecase
	sessionUsecase domain.SessionUsecase
	userUsecase    domain.UserUsecase
}

func NewAuthUsecase(
	googleUsecase domain.GoogleUsecase,
	sessionUsecase domain.SessionUsecase,
	userUsecase domain.UserUsecase,
) domain.AuthUsecase {
	return &authUsecase{
		googleUsecase:  googleUsecase,
		sessionUsecase: sessionUsecase,
		userUsecase:    userUsecase,
	}
}

func (u *authUsecase) Login(body payload.Login, ipAddress string) (*domain.User, *fiber.Cookie, error) {
	user, err := u.userUsecase.FindByEmail(body.Email)
	if err != nil {
		return nil, nil, err
	} else if user == nil {
		return nil, nil, errors.New("user not found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, nil, errors.New("password incorrect")
	}

	cookie, err := u.sessionUsecase.Create(user.Id, ipAddress)
	if err != nil {
		return nil, nil, err
	}
	
	return user, cookie, nil
}

func (u *authUsecase) GoogleLogin(c *fiber.Ctx) (*domain.User, *fiber.Cookie, error) {
	token, err := u.googleUsecase.GetToken(c)
	if err != nil {
		return nil, nil, err
	}

	profile, err := u.googleUsecase.GetProfile(token.AccessToken)
	if err != nil {
		return nil, nil, err
	}

	user, err := u.userUsecase.FindByEmail(profile.Email)
	if err != nil {
		return nil, nil, err
	}

	if user == nil {
		user, err = u.userUsecase.CreateFromGoogle(profile)
		if err != nil {
			return nil, nil, err
		}
	}

	cookie, err := u.sessionUsecase.Create(user.Id, c.IP())
	if err != nil {
		return nil, nil, err
	}

	return user, cookie, nil
}

func (u *authUsecase) Logout(sid string) (*fiber.Cookie, error) {
	err := u.sessionUsecase.Delete(sid)
	if err != nil {
		return nil, err
	}

	cookie := &fiber.Cookie{Name: "ssid", Expires: time.Unix(0, 0)}
	return cookie, nil
}
