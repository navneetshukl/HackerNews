package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello I am Navneet Shukla")
		return c.JSON(fiber.Map{
			"name": "Navneet Shukla",
		})
	})

	app.Listen(":8080")

}
