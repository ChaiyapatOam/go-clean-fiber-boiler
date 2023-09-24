package main

import (
	"log"

	"github.com/chaiyapatoam/go-clean-fiber-boiler/db"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/internal/config"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	env := config.LoadEnv()

	db, err := db.ConnectDB(env.MYSQL_URI)
	if err != nil {
		log.Fatal("Cant Connect To Mysql : ", err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"message": "Hello World"})
	})

	route.MainRoute(app, db, env)
	log.Fatal(app.Listen(":5000"))
}
