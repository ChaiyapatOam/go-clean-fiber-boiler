package main

import (
	"log"

	"github.com/chaiyapatoam/go-clean-fiber-boiler/db"
	"github.com/chaiyapatoam/go-clean-fiber-boiler/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Cant Connect To Mysql : ", err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"message": "Hello World"})
	})

	route.MainRoute(app, db)
	log.Fatal(app.Listen(":5000"))
}
