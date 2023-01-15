package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/", "./my-app/build")

	r.NoRoute(func(c *gin.Context) {
		c.File("./my-app/build/index.html")
	})
	r.Run(":3000")
}
