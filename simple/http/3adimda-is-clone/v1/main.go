package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"strconv"
	"bytes"
	"html/template"
	"github.com/patrickmn/go-cache"	
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var c = cache.New(1 * time.Minute, 1 * time.Minute)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string
	Gmail        string
	Password     string
	ProfilePhoto string
	Admin        string
	Company		 string
	CompanyID	 string
	CreatedAt    time.Time
}

type Company struct {
	ID   				primitive.ObjectID `bson:"_id,omitempty"`
	Name 				string
	Workers 			int
	AdvertisementCount	int
	Followers			int
	Views				int
	CreatedAt    		time.Time
}

type Message struct {
	ID   		primitive.ObjectID `bson:"_id,omitempty"`
	Title		string
	Message 	string
	Author		string
	AuthorID	string
	To 			string
	ToID		string
	CreatedAt   time.Time
}

type Works struct {
	ID   	 	primitive.ObjectID `bson:"_id,omitempty"`
	CompanyID	string
	Company		string
	UserID		string
	Photo 		string
	Title		string
	Description string
	CreatedAt   time.Time		
}

type Advert struct {
	ID   	 	primitive.ObjectID `bson:"_id,omitempty"`
	Company		string
	CompanyID	string
	Title		string
	Description	string
	Minprice	int
	Maxprice	int
	CreatedAt   time.Time
}

type Reply struct {
	ID   	 	primitive.ObjectID `bson:"_id,omitempty"`
	Company		string
	CompanyID	string
	Title		string
	Description string
	Author		string
	AuthorID	string
	MessageID	string
	ToID		string
	CreatedAt   time.Time		
}
func main(){
	r := gin.Default()
	r.LoadHTMLGlob("src/*.tmpl")
	r.Static("/static", "./static/")
	r.Static("/upload", "./upload/")

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

	usercollection := client.Database("findjob").Collection("users")
	companycollection := client.Database("findjob").Collection("companys")
	messagecollection := client.Database("findjob").Collection("messages")
	workscollection := client.Database("findjob").Collection("works")
	advertcollection := client.Database("findjob").Collection("adverts")
	replycollection := client.Database("findjob").Collection("replys")

	fmt.Println(usercollection, companycollection, messagecollection, workscollection, advertcollection, replycollection)
	//PAGES
	//Anasayfa 
	r.GET("/", func(ctx *gin.Context) {
		var adverts []Advert
		var user User
	
		cur, err := advertcollection.Find(ctx, bson.M{})
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer cur.Close(ctx)
	
		for cur.Next(ctx) {
			var advert Advert
			err := cur.Decode(&advert)
			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			adverts = append(adverts, advert)
		}
	
		if err := cur.Err(); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	
		userCookie, err := ctx.Cookie("token")
		if err == nil {
			var foundUser User
			oid, err := primitive.ObjectIDFromHex(userCookie)
			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			err = usercollection.FindOne(ctx, bson.M{"_id": oid}).Decode(&foundUser)
			if err != nil {
				user = foundUser
			}
		}
	
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
			"adverts": adverts,
			"user": user,
		}); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	
		c.Set("index.tmpl", buf.String(), cache.DefaultExpiration)
		ctx.Header("Content-Type", "text/html")
		ctx.Writer.Write(buf.Bytes())
	})
	//Kayıt olma sayfası
	r.GET("/register", func(ctx *gin.Context) {
		tmpl, found := c.Get("register.tmpl")
		if found {
			tmplStr, _ := tmpl.(string)
			ctx.Header("Content-Type", "text/html")
			ctx.Writer.WriteString(tmplStr)
			return
		}

		t, err := template.ParseFiles("src/register.tmpl")
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		buf := new(bytes.Buffer)
		if err := t.Execute(buf, gin.H{
			"title":  "Kayıt ol",
		}); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Set("register.tmpl", buf.String(), cache.DefaultExpiration)
		ctx.Header("Content-Type", "text/html")
		ctx.Writer.Write(buf.Bytes())
	})
	//Giriş sayfası
	r.GET("/login", func(ctx *gin.Context) {
		tmpl, found := c.Get("login.tmpl")
		if found {
			tmplStr, _ := tmpl.(string)
			ctx.Header("Content-Type", "text/html")
			ctx.Writer.WriteString(tmplStr)
			return
		}

		t, err := template.ParseFiles("src/login.tmpl")
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		buf := new(bytes.Buffer)
		if err := t.Execute(buf, gin.H{
			"title":  "Giriş yap",
		}); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Set("login.tmpl", buf.String(), cache.DefaultExpiration)
		ctx.Header("Content-Type", "text/html")
		ctx.Writer.Write(buf.Bytes())
	})
	//Kullanıcı paneli
	r.GET("/user/dashboard", func(ctx *gin.Context){
		cookie, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Hata": "Kayıt ol yada giriş yap",
			})
			return
		}
		objectID, err := primitive.ObjectIDFromHex(cookie)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Hata": "Kayıt ol yada giriş yap",
			})
			return
		}
		var user User
		userFilter := bson.M{"_id": objectID}
		err = usercollection.FindOne(ctx, userFilter).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Hata": "Kullanıcı bulunamadı",
			})
			return
		}
		ctx.HTML(http.StatusOK, "userdasboard.tmpl", gin.H{
			"user": user,
			"title": user.username,
		})
	})
	//Şirket paneli
	r.GET("/company/dashboard/:companyID", func(ctx *gin.Context){
		cookie, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Hata": "Kayıt ol yada giriş yap",
			})
			return
		}
		objectID, err := primitive.ObjectIDFromHex(cookie)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Hata": "Kayıt ol yada giriş yap",
			})
			return
		}
		var user User
		userFilter := bson.M{"_id": objectID}
		err = usercollection.FindOne(ctx, userFilter).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Hata": "Kullanıcı bulunamadı",
			})
			return
		}
		
	})
	//İlan paneli
	r.GET("/advert/dashboard/:advertID", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Hata": "Kayıt ol yada giriş yap",
			})
			return
		}
		objectID, err := primitive.ObjectIDFromHex(cookie)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Hata": "Kayıt ol yada giriş yap",
			})
			return
		}
		var user User
		userFilter := bson.M{"_id": objectID}
		err = usercollection.FindOne(ctx, userFilter).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Hata": "Kullanıcı bulunamadı",
			})
			return
		}

		advertID := ctx.Param("advdertID")
		oid, _:= primitive.ObjectIDFromHex(advertID)
		var advert Advert
		advertFilter := bson.M{"_id": oid}
		err = advertcollection.FindOne(ctx, advertFilter).Decode(&advert)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Hata": "İlan bulunamadı",
			})
			return
		}
		if advert.CompanyID == user.CompanyID {
		var messages []Message
		messageFilter := bson.M{"ToID":advert.CompanyID}
		cur, err := messagecollection.Find(ctx, messageFilter)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Hata": "Mesaj bulunamadı",
			})
			return
		}

		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var message Message
			err := cur.Decode(&message)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{
					"Hata": "Mesaj bulunamadı",
				})
				return
			}
			messages = append(messages, message)
		}

		ctx.HTML(http.StatusOK, "advertpanel.tmpl", gin.H{
			"title":   "Dashboard",
			"user":    user,
			"advert":  advert,
			"message": messages,
		})
	  } else {
		ctx.Redirect(http.StatusFound, "/")
	  }
	})
	//POST
	r.POST("/register", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		gmail := ctx.PostForm("gmail")

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
			ctx.HTML(http.StatusOK, "register.tmpl", gin.H{
				"message": "HATA Kullanıcı adı kullanılıyor!",
			})
		}
		if finduserwithmail != nil {
			ctx.HTML(http.StatusOK, "register.tmpl", gin.H{
				"message": "HATA Birisi Bu mail ile kayıt olmuş!",
			})
		}
		if err != nil {
			ctx.HTML(http.StatusOK, "register.tmpl", gin.H{
				"message": "HATA, Kayıt Başarısız.",
			})
		}

		ctx.SetCookie("token", insertResult.InsertedID.(primitive.ObjectID).Hex(), 3600, "/", "", false, true)
		ctx.Redirect(http.StatusFound, "/")
	})
	//Giriş yapma
	r.POST("/login", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		var user User
		err := usercollection.FindOne(context.TODO(), bson.M{"username": username, "password": password}).Decode(&user)
		if err != nil {
			ctx.HTML(http.StatusOK, "login.tmpl", gin.H{
				"message": "Kullanıcı adı veya şifre hatalı.",
			})
			return
		}

		ctx.SetCookie("token", user.ID.Hex(), 3600, "/", "", false, true)
		ctx.Redirect(http.StatusFound, "/")
	})
	//Şirket oluşturma
	r.POST("/create/company", func(ctx *gin.Context){
		userID, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Hata": "Giriş yap yada kayıt ol"})
			return
		}

		var user User
		userIDHex, _ := primitive.ObjectIDFromHex(userID)
		err = usercollection.FindOne(ctx, bson.M{"_id": userIDHex}).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Hata": "Giriş yap yada kayıt ol"})
			return
		}
	
		if user.CompanyID != "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"Hata": "Senin zaten şirketin var"})
			return
		}

		name := ctx.PostForm("name")
		if name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"Hata": "isim girilmemiş"})
			return
		}
	
		company := Company{
			ID:        primitive.NewObjectID(),
			Name:      name,
			Workers:   0,
			Followers: 0,
			Views:     0,
			CreatedAt: time.Now(),
		}
	
		_, err = companycollection.InsertOne(ctx, company)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Hata": "Şirket oluşturulamadı"})
			return
		}
	
		_, err = usercollection.UpdateOne(ctx, bson.M{"_id": userIDHex}, bson.M{"$set": bson.M{"companyid": company.ID.Hex()}},)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Hata": "Kullanıcı Güncellenemedi"})
			return
		}
	
		ctx.JSON(http.StatusOK, gin.H{"Mesaj": "Şirket Başarıyla oluşturuldu"})
	})
	//İlan paylaşma
	r.POST("/create/advert", func(ctx *gin.Context) {
		var advert Advert
		cookie, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Hata": "Kayıt ol yada giriş yap",
			})
			return
		}
		objectID, err := primitive.ObjectIDFromHex(cookie)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Hata": "Kayıt ol yada giriş yap",
			})
			return
		}
		var user User
		userFilter := bson.M{"_id": objectID}
		err = usercollection.FindOne(ctx, userFilter).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Hata": "Kullanıcı bulunamadı",
			})
			return
		}
		minprice, err := strconv.Atoi(ctx.PostForm("minprice"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Hata": "Geçersiz fiyat (en düşük)"})
			return
		}
		maxprice, err := strconv.Atoi(ctx.PostForm("maxprice"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Hata": "Geçersiz fiyat (en büyük)"})
			return
		}
		advert = Advert {
			ID: 			primitive.NewObjectID(),
			Title: 			ctx.PostForm("title"), 
			Description: 	ctx.PostForm("description"),
			CompanyID:		user.CompanyID,
			Minprice: 		minprice,
			Maxprice: 		maxprice,
			CreatedAt: 		time.Now(),
		}
		result, err := advertcollection.InsertOne(ctx, advert)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		fmt.Println(result)
		ctx.Redirect(http.StatusFound, "/")
	})
	//İlana başvurma
	r.POST("/apply/advert/:advertID", func(ctx *gin.Context){
		advertID:= ctx.Param("advertID")
		var message Message
		cookie, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Hata": "Kayıt ol yada giriş yap",
			})
			return
		}
		objectID, err := primitive.ObjectIDFromHex(cookie)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Hata": "Kayıt ol yada giriş yap",
			})
			return
		}
		var user User
		userFilter := bson.M{"_id": objectID}
		err = usercollection.FindOne(ctx, userFilter).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Hata": "Kullanıcı bulunamadı",
			})
			return
		}
		message = Message {
			ID: 			primitive.NewObjectID(),
			Title: 			ctx.PostForm("title"), 
			Message: 		ctx.PostForm("description"),
			Author: 		user.Username,
			AuthorID:		user.ID.Hex(),
			ToID:			advertID,		
			CreatedAt: 		time.Now(),
		}
		result, err := messagecollection.InsertOne(ctx, message)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		fmt.Println(result)
		ctx.Redirect(http.StatusFound, "/")
	})
	//Başvuruya Yanıt verme
	r.POST("/reply/message/:messageID", func(ctx *gin.Context){
		messageID:= ctx.Param("messageID")
		var reply Reply
		cookie, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Hata": "Kayıt ol yada giriş yap",
			})
			return
		}
		objectID, err := primitive.ObjectIDFromHex(cookie)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Hata": "Kayıt ol yada giriş yap",
			})
			return
		}
		var user User
		userFilter := bson.M{"_id": objectID}
		err = usercollection.FindOne(ctx, userFilter).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Hata": "Kullanıcı bulunamadı",
			})
			return
		}
		oid, _ := primitive.ObjectIDFromHex(messageID)
		
		var message Message
		messageFilter := bson.M{"_id": oid}
		err = messagecollection.FindOne(ctx, messageFilter).Decode(&message)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Hata": "mesaj bulunamadı",
			})
			return
		}
		reply = Reply {
			ID: 			primitive.NewObjectID(),
			Title: 			ctx.PostForm("title"), 
			Description: 	ctx.PostForm("description"),
			Author: 		user.Username,
			AuthorID:		user.ID.Hex(),
			MessageID:		messageID,	
			ToID:			message.AuthorID,	
			CreatedAt: 		time.Now(),
		}
		result, err := replycollection.InsertOne(ctx, reply)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		fmt.Println(result)
		ctx.Redirect(http.StatusFound, "/")
	})

	r.Run(":5000")
}