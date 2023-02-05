package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx){
	  c.SendFile("index.html")
	}
	app.Post("/giris", func(c *fiber.Ctx) {
	  isim := c.FormValue("isim")
	  soyisim := c.FormValue("soyisim")
	  c.Send(isim,soyisim)
	})

	log.Fatal(app.Listen(":3000"))
}
