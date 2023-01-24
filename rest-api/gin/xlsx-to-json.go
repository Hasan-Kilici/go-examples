package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

func main() {
	r := gin.Default()
	r.Static("/", "./my-app/build")
	r.POST("/api/products", func(c *gin.Context) {
		file, err := xlsx.OpenFile("urun.xlsx")
		if err != nil {
			panic(err)
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
			panic(err)
		}

		fmt.Println(string(data))
		c.JSON(http.StatusOK, string(data))
	})
	r.Run(":8080")
}
