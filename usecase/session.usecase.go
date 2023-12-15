package usecase

import (
	"crypto/md5"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"strings"
	"time"

	"github.com/chaiyapatoam/go-clean-fiber-boiler/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type sessionUsecase struct {
	env               *domain.Env
	sessionRepository domain.SessionRepository
}

func NewSessionUsecase(env *domain.Env, sessionRepository domain.SessionRepository) domain.SessionUsecase {
	return &sessionUsecase{
		env:               env,
		sessionRepository: sessionRepository,
	}
}

func (u *sessionUsecase) Sign(id string) string {
	h := md5.New()
	h.Write([]byte(u.env.SESSION_SECRET))
	sig1 := h.Sum(nil)
	h.Write([]byte(id))
	sig2 := h.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(sig1) + base64.StdEncoding.EncodeToString(sig2)
	return u.env.SESSION_PREFIX + ":" + id + "." + signature
}

func (u *sessionUsecase) Unsign(header string) (string, error) {
	prefix := u.env.SESSION_PREFIX
	if !strings.HasPrefix(header, prefix+":") {
		return "", errors.New("prefix mismatch")
	}

	id := header[len(prefix)+1 : strings.LastIndex(header, ".")]
	expectation := u.Sign(id)

	isLengthMatch := len([]byte(header)) == len([]byte(expectation))
	isInputMatch := subtle.ConstantTimeCompare([]byte(header), []byte(expectation)) == 1

	if !isLengthMatch || !isInputMatch {
		return "", errors.New("signature mismatch")
	}
	return id, nil
}

func (u *sessionUsecase) Create(userId string, ipAddress string) (*fiber.Cookie, error) {
	id := uuid.NewString()
	createdAt := time.Now()
	expiresAt := createdAt.Add(3600 * time.Second * 24) // 24hr

	err := u.sessionRepository.Create(&domain.Session{
		Id:        id,
		UserId:    userId,
		IpAddress: ipAddress,
		CreatedAt: createdAt,
		ExpiredAt: expiresAt,
	})

	if err != nil {
		return nil, err
	}

	signId := u.Sign(id)

	cookie := &fiber.Cookie{
		Name:     "ssid",
		Value:    signId,
		HTTPOnly: true,
		Secure:   true,
		Expires:  expiresAt,
	}
	return cookie, nil
}

func (u *sessionUsecase) Get(ssid string) (*domain.Session, error) {
	id, err := u.Unsign(ssid)
	if err != nil {
		return nil, err
	}
	
	session, err := u.sessionRepository.Get(id)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (u *sessionUsecase) Delete(ssid string) error {
	err := u.sessionRepository.Delete(ssid)

	if err != nil {
		return err
	}
	return nil
}

func (u *sessionUsecase) Validate(ssid string) (bool, string, error) {
	session, err := u.Get(ssid)
	if session == nil {
		return false, "", errors.New("session not found")
	} else if err != nil {
		return false, "", err
	}

	// if Expired
	if !time.Now().Before(session.ExpiredAt) {
		return false, "", err
	}
	return true, session.UserId, nil
}
