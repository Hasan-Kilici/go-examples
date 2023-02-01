package main

import (
	"html/template"
	"net/http"
	"time"
  "bytes"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var c = cache.New(5*time.Minute, 10*time.Minute)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("src/*.tmpl")
	r.Static("/static", "./static/")

	r.GET("/", func(ctx *gin.Context) {
		tmpl, found := c.Get("index.tmpl")
		if found {
			tmplStr, _ := tmpl.(string)
			ctx.Header("Content-Type", "text/html")
			ctx.Writer.WriteString(tmplStr)
			return
		}

		t, err := template.ParseFiles("src/index.tmpl")
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		buf := new(bytes.Buffer)
		if err := t.Execute(buf, gin.H{
			"title":  "Anasayfa",
			"author": "Hasan KILICI",
		}); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Set("index.tmpl", buf.String(), cache.DefaultExpiration)
		ctx.Header("Content-Type", "text/html")
		ctx.Writer.Write(buf.Bytes())
	})

	r.GET("/about", func(ctx *gin.Context) {
		about, found := c.Get("about")
		if found {
			aboutBytes, _ := about.([]byte)
			ctx.Data(http.StatusOK, "text/html", aboutBytes)
			return
		}

		t, err := template.ParseFiles("src/about.tmpl")
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		buf := new(bytes.Buffer)
		if err := t.Execute(buf, gin.H{
			"title":  "Hakkımda",
			"author": "Hasan KILICI",
		}); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Set("about", buf.Bytes(), cache.DefaultExpiration)
		ctx.Data(http.StatusOK, "text/html", buf.Bytes())
	})

	r.Run()
}
