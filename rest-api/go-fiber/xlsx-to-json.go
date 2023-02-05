package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/tealeg/xlsx"
)

func main() {
	app := fiber.New()

	app.Post("/api/products", func(c *fiber.Ctx) {
		file, err := xlsx.OpenFile("urun.xlsx")
		if err != nil {
			c.Status(500).Send(err)
			return
		}

		sheet := file.Sheets[0]
		rows := make([]map[string]string, 0)
		for _, row := range sheet.Rows {
			rowMap := make(map[string]string)
			for j, cell := range row.Cells {
				rowMap[sheet.Rows[0].Cells[j].String()] = cell.String()
			}
			rows = append(rows, rowMap)
		}

		data, err := json.Marshal(rows)
		if err != nil {
			c.Status(500).Send(err)
			return
		}

		fmt.Println(string(data))
		c.JSON(data)
	})

	app.Listen(8080)
}
