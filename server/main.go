package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var requests = make(map[string]int)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")

	api.Get("/send", func(c *fiber.Ctx) error {
		queries := c.Queries()
		requests[queries["client"]] += 1
		return c.SendString(fmt.Sprint(requests))
	})

	api.Get("/get", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprint(requests))
	})

	log.Fatal(app.Listen(":5000"))
}
