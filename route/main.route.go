package route

import (
	"github.com/chaiyapatoam/go-clean-fiber-boiler/controller"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/repository"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func MainRoute(app *fiber.App, db *sqlx.DB) {
	userRepository := repository.NewuserRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)

	userController := controller.NewUserController(userUsecase)

	users := app.Group("/api/users")
	users.Get("/:id", userController.Get)
	users.Post("/", userController.Create)

}
