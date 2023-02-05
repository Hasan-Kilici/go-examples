package main

import (
	"fmt"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.SendFile("index.html")
	})

	app.Post("/upload", func(c *fiber.Ctx) {
		file, err := c.FormFile("file")
		if err != nil {
			c.Status(400).SendString(err.Error())
			return
		}
		fmt.Println(file.Filename)
		dst := "uploads/" + file.Filename
		if err := c.SaveFile(file, dst); err != nil {
			c.Status(400).SendString(err.Error())
			return
		}
		c.SendString(fmt.Sprintf("'%s' yüklendi!", file.Filename))
	})

	app.Listen(3000)
}
