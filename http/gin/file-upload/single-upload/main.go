package main

import (
  "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
  r.MaxMultipartMemory = 8 << 20  
	r.LoadHTMLFiles("index.html")
  
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "anasayfa",
		})
	})

  r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)
    dst := "uploads/" + file.Filename
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' yüklendi!", file.Filename))
	})
  
	r.Run()
}
