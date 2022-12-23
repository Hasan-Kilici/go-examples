package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.LoadHTMLFiles("index.html");
  r.GET("/", func(c *gin.Context) {
   c.HTML(http.StatusOK, "index.html",gin.H{
			"content": "This is an index page...",
		});
  })
  r.Run()
}
