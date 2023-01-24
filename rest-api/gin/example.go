package main

import (
    "github.com/gin-gonic/gin"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    r := gin.Default()

    r.GET("/users", func(c *gin.Context) {
        users := []User{
            {Name: "John", Age: 30},
            {Name: "Mike", Age: 45},
        }
        c.JSON(200, users)
    })

    r.Run(":8080")
}
