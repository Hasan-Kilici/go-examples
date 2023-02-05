package main

import (
	"encoding/json"
	"github.com/gofiber/fiber"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	app := fiber.New()

	app.Get("/users", func(c *fiber.Ctx) {
		users := []User{
			{Name: "John", Age: 30},
			{Name: "Mike", Age: 45},
		}

		json, _ := json.Marshal(users)

		c.Type("application/json")
		c.Send(json)
	})

	app.Listen(8080)
}
