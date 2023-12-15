package controller

import (
	"fmt"

	"github.com/chaiyapatoam/go-clean-fiber-boiler/domain"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/payload"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/response"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/validator"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	env            *domain.Env
	authUsecase    domain.AuthUsecase
	sessionUsecase domain.SessionUsecase
	googleUsecase  domain.GoogleUsecase
	userUsecase    domain.UserUsecase
}

func NewAuthController(env *domain.Env, authUsecase domain.AuthUsecase, sessionUsecase domain.SessionUsecase, googleUsecase domain.GoogleUsecase, userUsecase domain.UserUsecase) *AuthController {
	return &AuthController{
		env:            env,
		authUsecase:    authUsecase,
		sessionUsecase: sessionUsecase,
		googleUsecase:  googleUsecase,
		userUsecase:    userUsecase}
}

func (auth *AuthController) GetGoogleUrl(c *fiber.Ctx) error {
	path := auth.googleUsecase.GoogleConfig()
	url := path.AuthCodeURL("state")

	return c.Status(200).JSON(fiber.Map{"success": true, "url": url})
}

func (auth *AuthController) SignUp(c *fiber.Ctx) error {
	var body payload.Register
	err := validator.ValidateBody(c, &body)
	if err != nil {
		return response.ErrorResponse(c, 400, err.Error())
	}

	err = auth.userUsecase.Create(body)
	if err != nil {
		return err
	}

	return response.SuccessResponse(c, "Created")
}

func (auth *AuthController) SignIn(c *fiber.Ctx) error {
	var body payload.Login
	err := validator.ValidateBody(c, &body)
	if err != nil {
		return response.ErrorResponse(c, 400, err.Error())
	}

	ipAddress := c.IP()

	user, cookie, err := auth.authUsecase.Login(body, ipAddress)
	if err != nil {
		return err
	}

	if user == nil {
		return response.ErrorResponse(c, 500, "")
	}

	c.Cookie(cookie)
	return response.SuccessResponse(c, "Login Success")
}

func (auth *AuthController) SignInWithGoogle(c *fiber.Ctx) error {
	user, cookie, err := auth.authUsecase.GoogleLogin(c)
	if err != nil {
		fmt.Println("SignInWithGoogle Error: ", err)
		return c.Redirect(auth.env.FRONTEND_URL)
	}

	if user == nil {
		return response.ErrorResponse(c, 500, "")
	}

	c.Cookie(cookie)
	return c.Redirect(auth.env.FRONTEND_URL)
}

func (auth *AuthController) Logout(c *fiber.Ctx) error {
	ssid := c.Cookies("ssid")
	cookie, err := auth.authUsecase.Logout(ssid)

	if err != nil {
		return err
	}

	c.Cookie(cookie)
	return response.SuccessResponse(c, "Logout Success")
}
