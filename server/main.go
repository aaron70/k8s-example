package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var m sync.Mutex
var requests int = 0

func main() {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")

	api.Get("/send", func(c *fiber.Ctx) error {
		m.Lock()
		requests += 1
		m.Unlock()
		return c.SendString(fmt.Sprintf("Cantidad de requests: %d", requests))
	})

	api.Get("/get", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Cantidad de requests: %d", requests))
	})

	log.Fatal(app.Listen(":5000"))
}
