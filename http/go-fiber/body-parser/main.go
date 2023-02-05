package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Post("/form", func(c *fiber.Ctx) {
		isim := c.FormValue("isim")
	  soyisim := c.FormValue("soyisim")

		header := fiber.Map{
			"isim": isim,
      "soyisim": soyisim,
		}

		c.Render("index.html", header)
	})

	log.Fatal(app.Listen(":3000"))
}
