package controller

import (
	"errors"

	"github.com/chaiyapatoam/go-clean-fiber-boiler/domain"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/payload"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/response"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/validator"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/middleware"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (userController *UserController) Get(c *fiber.Ctx) error {
	userId := middleware.GetUser(c)
	user, err := userController.userUsecase.Get(userId)

	if err != nil {
		return errors.New("error get user")
	}

	return response.SuccessDataResponse(c, user)
}

func (userController *UserController) ChangePassword(c *fiber.Ctx) error {
	var body payload.ChangePassword
	err := validator.ValidateBody(c, &body)
	if err != nil {
		return response.ErrorResponse(c, 400, err.Error())
	}

	userId := middleware.GetUser(c)
	err = userController.userUsecase.ChangePassword(userId, body)

	if err != nil {
		return err
	}

	return response.SuccessResponse(c, "")
}
