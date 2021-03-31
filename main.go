package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	type Post struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Rev    uint8  `json:"rev"`
	}

	dataReturn := []Post{
		{
			Title:  "I am the best",
			Author: "Teerapat",
			Rev:    1,
		},
		{
			Title:  "I am the worst",
			Author: "Xeus",
			Rev:    2,
		},
		{
			Title:  "Freaking Stupid",
			Author: "Prommarak",
			Rev:    4,
		},
		{
			Title:  "I am the Thanos",
			Author: "Thanos",
			Rev:    3,
		},
	}

	app.Get("/thanos", func(c *fiber.Ctx) error {
		return c.JSON(dataReturn)
	})

	app.Listen(":3000")
}
