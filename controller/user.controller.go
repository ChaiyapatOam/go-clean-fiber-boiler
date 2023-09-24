package controller

import (
	"github.com/chaiyapatoam/go-clean-fiber-boiler/domain"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/payload"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/response"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/validator"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (userController *UserController) Create(c *fiber.Ctx) error {
	var body payload.CreateUser
	err := validator.ValidateBody(c, &body)
	if err != nil {
		return response.ErrorResponse(c, 400, err.Error())
	}

	err = userController.userUsecase.Create(body)
	if err != nil {
		return err
	}

	return response.SuccessResponse(c, "Created")
}

func (userController *UserController) Get(c *fiber.Ctx) error {
	user, err := userController.userUsecase.Get(c.Params("id"))

	if err != nil {
		return err
	}
	return response.SuccessDataResponse(c,user)
}
