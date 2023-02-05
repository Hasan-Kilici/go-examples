package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

func main() {
	app := fiber.New()

	app.Settings.Templates = html.New("./templates", ".html")

	app.Get("/", func(c *fiber.Ctx) {
		if err := c.Render("index", fiber.Map{
			"title": "anasayfa",
		}); err != nil {
			c.Status(500).Send(err)
		}
	})

	app.Post("/upload", func(c *fiber.Ctx) {
		files, err := c.FormFileAll("upload[]")
		if err != nil {
			c.Status(500).Send(err)
			return
		}

		for _, file := range files {
			fmt.Println(file.Filename)
			if err := file.Save("./uploads/" + file.Filename); err != nil {
				c.Status(500).Send(err)
				return
			}
		}
		c.Send(fmt.Sprintf("%d files uploaded!", len(files)))
	})

	app.Listen(":3000")
}
