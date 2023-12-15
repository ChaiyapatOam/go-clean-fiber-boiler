package route

import (
	"github.com/chaiyapatoam/go-clean-fiber-boiler/controller"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/domain"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/middleware"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/repository"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func MainRoute(app *fiber.App, db *sqlx.DB, env *domain.Env) {
	userRepository := repository.NewuserRepository(db)
	sessionRepository := repository.NewSessionRepository(db)

	googleUsecase := usecase.NewGoogleUsecase(env)
	sessionUsecase := usecase.NewSessionUsecase(env, sessionRepository)
	userUsecase := usecase.NewUserUsecase(userRepository, sessionUsecase)
	authUsecase := usecase.NewAuthUsecase(googleUsecase, sessionUsecase, userUsecase)

	userController := controller.NewUserController(userUsecase)
	authController := controller.NewAuthController(env, authUsecase, sessionUsecase, googleUsecase, userUsecase)

	authMiddleWare := middleware.NewAuthMiddleWare(sessionUsecase)

	auth := app.Group("/api/auth")
	auth.Post("/login", authMiddleWare, authController.SignIn)
	auth.Post("/register", authController.SignUp)
	auth.Get("/google", authController.GetGoogleUrl)
	auth.Get("/google/callback", authController.SignInWithGoogle)
	auth.Post("/logout", authMiddleWare, authController.Logout)

	users := app.Group("/api/users")
	users.Get("/", authMiddleWare, userController.Get)
	// users.Patch("/",authMiddleWare,userController.Update)
	users.Patch("/change-password", authMiddleWare, userController.ChangePassword)

}
