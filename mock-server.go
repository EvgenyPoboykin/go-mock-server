package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func handler(c *fiber.Ctx) error {
	return c.SendString("I'm a POST request!")
}

func main() {
	app := fiber.New()

	app.Post("/mock", handler)

	log.Fatal(app.Listen(":3000"))
}
