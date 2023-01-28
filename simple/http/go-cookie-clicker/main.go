package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/martinlindhe/notify"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("src/*.tmpl")
	r.Static("/static", "./static/")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Home",
		})
	})

	r.GET("/shop", func(c *gin.Context) {
		c.HTML(http.StatusOK, "shop.tmpl", func(c *gin.Context) {

		})
	})

	r.POST("/buy/cursor", func(c *gin.Context) {
		cash, err := c.Cookie("cash")
		cursor, _ := c.Cookie("cursor")
		perSecond, _ := c.Cookie("persecond")

		usercash, _ := strconv.ParseInt(cash, 10, 64)
		usercursor, _ := strconv.ParseInt(cursor, 10, 64)
		userPerSecond, _ := strconv.ParseInt(perSecond, 10, 64)

		var fiyat int64 = 100
		if usercash > fiyat {
			notify.Notify("Cookie Clicker", "EŞYA BAŞARIYLA SATIN ALINDI", "", "")
			usercash = usercash - fiyat
			nowcash := strconv.FormatInt(usercash, 10)
			usercursor = usercursor + 1
			nowcursor := strconv.FormatInt(usercursor, 10)
			userPerSecond = userPerSecond + 1
			nowPerSecond := strconv.FormatInt(userPerSecond, 10)

			c.SetCookie("cash", nowcash, 60*60*24, "/", "https://Gin-Go-5.hasan-kilici.repl.co", true, true)
			c.SetCookie("cursor", nowcursor, 60*60*24, "/", "https://Gin-Go-5.hasan-kilici.repl.co", true, true)
			c.SetCookie("persecond", nowPerSecond, 60*60*24, "/", "https://Gin-Go-5.hasan-kilici.repl.co", true, true)

			fmt.Println("Cursor satın alındı")
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "Home",
			})
		} else {
			fmt.Println("Kullanıcın Parası Bu ürünü Almaya yetmiyor.")
			c.HTML(http.StatusOK, "error.tmpl", gin.H{
				"title": "Hata",
			})
		}

		if err != nil {
			fmt.Println("Çerez Bulunamadı")
		}
	})

	r.POST("/buy/grandma", func(c *gin.Context) {
		cash, err := c.Cookie("cash")
		grandma, _ := c.Cookie("grandma")
		perSecond, _ := c.Cookie("persecond")

		usercash, _ := strconv.ParseInt(cash, 10, 64)
		usergrandma, _ := strconv.ParseInt(grandma, 10, 64)
		userPerSecond, _ := strconv.ParseInt(perSecond, 10, 64)

		var fiyat int64 = 1000
		if usercash > fiyat {
			notify.Notify("Cookie Clicker", "EŞYA BAŞARIYLA SATIN ALINDI", "", "")
			usercash = usercash - fiyat
			nowcash := strconv.FormatInt(usercash, 10)
			usergrandma = usergrandma + 1
			nowgrandma := strconv.FormatInt(usergrandma, 10)
			userPerSecond = userPerSecond + 5
			nowPerSecond := strconv.FormatInt(userPerSecond, 10)

			c.SetCookie("cash", nowcash, 60*60*24, "/", "https://Gin-Go-5.hasan-kilici.repl.co", true, true)
			c.SetCookie("grandma", nowgrandma, 60*60*24, "/", "https://Gin-Go-5.hasan-kilici.repl.co", true, true)
			c.SetCookie("persecond", nowPerSecond, 60*60*24, "/", "https://Gin-Go-5.hasan-kilici.repl.co", true, true)

			fmt.Println("Cursor satın alındı")
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "Home",
			})
		} else {
			fmt.Println("Kullanıcın Parası Bu ürünü Almaya yetmiyor.")
			c.HTML(http.StatusOK, "error.tmpl", gin.H{
				"title": "Hata",
			})
		}

		if err != nil {
			fmt.Println("Çerez Bulunamadı")
		}
	})

	r.POST("/buy/farm", func(c *gin.Context) {
		cash, err := c.Cookie("cash")
		farm, _ := c.Cookie("farm")
		perSecond, _ := c.Cookie("persecond")

		usercash, _ := strconv.ParseInt(cash, 10, 64)
		userfarm, _ := strconv.ParseInt(farm, 10, 64)
		userPerSecond, _ := strconv.ParseInt(perSecond, 10, 64)

		var fiyat int64 = 10000
		if usercash > fiyat {
			notify.Notify("Cookie Clicker", "EŞYA BAŞARIYLA SATIN ALINDI", "", "")
			usercash = usercash - fiyat
			nowcash := strconv.FormatInt(usercash, 10)
			userfarm = userfarm + 1
			nowfarm := strconv.FormatInt(userfarm, 10)
			userPerSecond = userPerSecond + 15
			nowPerSecond := strconv.FormatInt(userPerSecond, 10)

			c.SetCookie("cash", nowcash, 60*60*24, "/", "https://Gin-Go-5.hasan-kilici.repl.co", true, true)
			c.SetCookie("farm", nowfarm, 60*60*24, "/", "https://Gin-Go-5.hasan-kilici.repl.co", true, true)
			c.SetCookie("persecond", nowPerSecond, 60*60*24, "/", "https://Gin-Go-5.hasan-kilici.repl.co", true, true)

			fmt.Println("Cursor satın alındı")
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "Home",
			})
		} else {
			fmt.Println("Kullanıcın Parası Bu ürünü Almaya yetmiyor.")
		}

		if err != nil {
			fmt.Println("Çerez Bulunamadı")
		}
	})

	r.POST("/buy/factory", func(c *gin.Context) {
		cash, err := c.Cookie("cash")
		factory, _ := c.Cookie("factory")
		perSecond, _ := c.Cookie("persecond")

		usercash, _ := strconv.ParseInt(cash, 10, 64)
		userfactory, _ := strconv.ParseInt(factory, 10, 64)
		userPerSecond, _ := strconv.ParseInt(perSecond, 10, 64)

		var fiyat int64 = 50000
		if usercash > fiyat {
			notify.Notify("Cookie Clicker", "EŞYA BAŞARIYLA SATIN ALINDI", "", "")
			usercash = usercash - fiyat
			nowcash := strconv.FormatInt(usercash, 10)
			userfactory = userfactory + 1
			nowfactory := strconv.FormatInt(userfactory, 10)
			userPerSecond = userPerSecond + 30
			nowPerSecond := strconv.FormatInt(userPerSecond, 10)

			c.SetCookie("cash", nowcash, 60*60*24, "/", "https://Gin-Go-5.hasan-kilici.repl.co", true, true)
			c.SetCookie("factory", nowfactory, 60*60*24, "/", "https://Gin-Go-5.hasan-kilici.repl.co", true, true)
			c.SetCookie("persecond", nowPerSecond, 60*60*24, "/", "https://Gin-Go-5.hasan-kilici.repl.co", true, true)

			fmt.Println("Cursor satın alındı")
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "Home",
			})
		} else {
			fmt.Println("Kullanıcın Parası Bu ürünü Almaya yetmiyor.")
			c.HTML(http.StatusOK, "error.tmpl", gin.H{
				"title": "Hata",
			})
		}

		if err != nil {
			fmt.Println("Çerez Bulunamadı")
		}
	})

	r.Run(":4600")
}
