package main

import (
	"log"

	"github.com/chaiyapatoam/go-clean-fiber-boiler/db"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/config"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func main() {
	env := config.LoadEnv()

	db, err := db.ConnectDB(env.MYSQL_URI)
	if err != nil {
		log.Fatal("Cant Connect To Mysql : ", err)
	}

	app := fiber.New()
	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "http://localhost:3000, http://localhost:5173",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"message": "Hello World"})
	})

	route.MainRoute(app, db, env)
	log.Fatal(app.Listen(":" + env.PORT))
}
