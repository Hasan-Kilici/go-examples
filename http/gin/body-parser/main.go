package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "anasayfa",
		})
	})

	r.POST("/giris", func(c *gin.Context) {
		isim := c.PostForm("isim")
		soyisim := c.PostForm("soyisim")
		c.String(http.StatusOK, "%s %s", isim, soyisim)
	})
	r.Run()
}
