package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string
	Gmail        string
	Password     string
	ProfilePhoto string
	Admin        string
	Views        int
	Likes        int
	CreatedAt    time.Time
}

type Blog struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string
	Description string
	Banner      string
	Html        string
	Author      string
}

type New struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string
	Description string
	Banner      string
	Html        string
	Author      string
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodburi")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB bağlantısı başarılı!")

	usercollection := client.Database("newswebsite").Collection("users")
	blogcollection := client.Database("newswebsite").Collection("blog")
	newcollection := client.Database("newswebsite").Collection("news")

	r := gin.Default()
	r.LoadHTMLGlob("src/*.tmpl")
	r.Static("/static", "./static/")
	r.Static("/upload", "./upload/")

	//Sayfalar
	//Anasayfa
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Anasayfa",
		})
	})
	//Giriş yapma sayfası
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title": "Giriş yap",
		})
	})
	//Kayıt olma sayfası
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.tmpl", gin.H{
			"title": "Kayıt ol",
		})
	})
	r.GET("/new/:id", func(c *gin.Context) {
		id := c.Param("id")
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
			return
		}
		var news New
		err = newcollection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&news)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Haber bulunamadı"})
			return
		}
		c.HTML(http.StatusOK, "news.tmpl", gin.H{
			"title":       news.Title,
			"description": news.Description,
			"banner":      news.Banner,
			"html":        news.Html,
			"author":      news.Author,
		})
	})

	// Blog sayfası
	r.GET("/blog/:id", func(c *gin.Context) {
		id := c.Param("id")
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
			return
		}
		var blog Blog
		err = blogcollection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&blog)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog bulunamadı"})
			return
		}
		c.HTML(http.StatusOK, "blog.tmpl", gin.H{
			"title":       blog.Title,
			"description": blog.Description,
			"banner":      blog.Banner,
			"html":        blog.Html,
			"author":      blog.Author,
		})
	})
	//Dashboard sayfası
	r.GET("/dashboard", func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			fmt.Println("İzinsiz giriş yapıldı!")
			c.Redirect(http.StatusFound, "/")
			return
		}
		oid, err2 := primitive.ObjectIDFromHex(token)
		if err2 != nil {
			fmt.Println(err2)
			fmt.Println("ID Mevcut değil.")
			return
		}
		var user User
		err2 = usercollection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&user)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		fmt.Println("Kullanıcı:", user)
		if user.Admin == "True" {
			c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{
				"title": "Admin dashboard",
			})
		}
	})
	//POST
	//Kayıt olma
	r.POST("/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		gmail := c.PostForm("gmail")

		fmt.Println(username, password, gmail)

		user := User{
			ID:           primitive.NewObjectID(),
			Username:     username,
			Password:     password,
			ProfilePhoto: "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png",
			Gmail:        gmail,
			Admin:        "False",
			CreatedAt:    time.Now(),
		}

		insertResult, err := usercollection.InsertOne(context.TODO(), user)
		finduserwithmail := usercollection.FindOne(context.TODO(), bson.M{"gmail": gmail}).Decode(&user)
		finduserwithusername := usercollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
		if finduserwithusername != nil {
			c.HTML(http.StatusOK, "register.tmpl", gin.H{
				"message": "HATA Kullanıcı adı kullanılıyor!",
			})
		}
		if finduserwithmail != nil {
			c.HTML(http.StatusOK, "register.tmpl", gin.H{
				"message": "HATA Birisi Bu mail ile kayıt olmuş!",
			})
		}
		if err != nil {
			c.HTML(http.StatusOK, "register.tmpl", gin.H{
				"message": "HATA, Kayıt Başarısız.",
			})
		}

		c.SetCookie("token", insertResult.InsertedID.(primitive.ObjectID).Hex(), 3600, "/", "", false, true)
		c.Redirect(http.StatusFound, "/")
	})
	//Giriş yapma
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		var user User
		err := usercollection.FindOne(context.TODO(), bson.M{"username": username, "password": password}).Decode(&user)
		if err != nil {
			c.HTML(http.StatusOK, "login.tmpl", gin.H{
				"message": "Kullanıcı adı veya şifre hatalı.",
			})
			return
		}

		c.SetCookie("token", user.ID.Hex(), 3600, "/", "", false, true)
		c.Redirect(http.StatusFound, "/")
	})
	//Blog Yükleme
	r.POST("/add/blog", func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			fmt.Println("izinsiz giriş!")
			c.Redirect(http.StatusFound, "/")
			return
		}
		oid, err2 := primitive.ObjectIDFromHex(token)
		if err2 != nil {
			fmt.Println(err2)
			fmt.Println("ID Mevcut değil.")
			return
		}
		var user User
		err2 = usercollection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&user)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		fmt.Println("Kullanıcı:", user)
		if user.Admin == "True" {
			btitle := c.PostForm("btitle")
			bdescription := c.PostForm("bdescription")
			bhtml := c.PostForm("bhtml")
			bauthor := user.Username

			file, _ := c.FormFile("file")
			fmt.Println(file.Filename)
			dst := "upload/" + file.Filename

			blog := Blog{
				Title:       btitle,
				Description: bdescription,
				Html:        bhtml,
				Author:      bauthor,
				Banner:      dst,
			}

			c.SaveUploadedFile(file, dst)

			insertResult, mongoerr := blogcollection.InsertOne(context.TODO(), blog)
			if mongoerr != nil {
				fmt.Println("MongoDB hatası...")
			}
			fmt.Println(insertResult)
			c.Redirect(http.StatusFound, "/dashboard")
		} else {
			fmt.Println("Kullanıcı Admin değil!")
		}
	})
	//Haber Yükleme
	r.POST("/add/new", func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			fmt.Println("izinsiz giriş!")
			c.Redirect(http.StatusFound, "/")
			return
		}
		oid, err2 := primitive.ObjectIDFromHex(token)
		if err2 != nil {
			fmt.Println(err2)
			fmt.Println("ID Mevcut değil.")
			return
		}
		var user User
		err2 = usercollection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&user)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		fmt.Println("Kullanıcı:", user)
		if user.Admin == "True" {
			ntitle := c.PostForm("ntitle")
			ndescription := c.PostForm("ndescription")
			nhtml := c.PostForm("nhtml")
			nauthor := user.Username

			file, _ := c.FormFile("nfile")
			fmt.Println(file.Filename)
			dst := "upload/" + file.Filename

			new := New{
				Title:       ntitle,
				Description: ndescription,
				Html:        nhtml,
				Author:      nauthor,
				Banner:      dst,
			}

			c.SaveUploadedFile(file, dst)

			insertResult, mongoerr := newcollection.InsertOne(context.TODO(), new)
			if mongoerr != nil {
				fmt.Println("MongoDB hatası...")
			}
			fmt.Println(insertResult)
			c.Redirect(http.StatusFound, "/dashboard")
		} else {
			fmt.Println("Kullanıcı Admin değil!")
		}
	})
	r.Run(":5000")
}
