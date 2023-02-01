package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}

type Task struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserId string
	Task   string
	Status string
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

	usercollection := client.Database("todolist").Collection("users")
	taskcollection := client.Database("todolist").Collection("tasks")

	r := gin.Default()
	r.LoadHTMLGlob("pages/*.tmpl")
	r.Static("/static", "./static/")
	//GET
	r.GET("/", func(c *gin.Context) {
		usertoken, err := c.Cookie("token")
		if err != nil {
			user := User{
				ID: primitive.NewObjectID(),
			}
			insertResult, err := usercollection.InsertOne(context.TODO(), user)
			if err != nil {
				log.Fatal(err)
			}
			c.SetCookie("token", insertResult.InsertedID.(primitive.ObjectID).Hex(), 3600, "/", "", false, true)
			c.HTML(http.StatusOK, "todo.tmpl", gin.H{
				"message": fmt.Sprintf("Kayıt başarılı. Kullanıcı ID'niz: %s", insertResult.InsertedID),
				"title":   "Todo List",
				"userID":  usertoken,
			})
		} else {
			// task'leri bulun
			var tasks []Task
			cur, err := taskcollection.Find(context.TODO(), bson.M{"userid": usertoken})
			if err != nil {
				log.Fatalf("Error finding task: %v", err)
			}
			defer cur.Close(context.TODO())
			for cur.Next(context.TODO()) {
				var task Task
				err := cur.Decode(&task)
				if err != nil {
					log.Fatalf("Error decoding task: %v", err)
				}
				tasks = append(tasks, task)
			}
			if err := cur.Err(); err != nil {
				log.Fatalf("Error iterating tasks: %v", err)
			}

			c.HTML(http.StatusOK, "todo.tmpl", gin.H{
				"title": "Todo List",
				"tasks": tasks,
			})
			fmt.Println(usertoken)
		}
	})
	//POST
	//Add Task
	r.POST("/add-task", func(c *gin.Context) {
		task := c.PostForm("task")
		user, err := c.Cookie("token")
		fmt.Println(task)
		if err != nil {
			log.Fatal(err)
		} else {
			tasks := Task{
				ID:     primitive.NewObjectID(),
				UserId: user,
				Task:   task,
				Status: "not-finished",
			}
			insertResult, err := taskcollection.InsertOne(context.TODO(), tasks)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(insertResult)
				if err != nil {
					log.Fatal(err)
				} else {
					c.Redirect(http.StatusFound, "/")
				}
			}
		}
	})
	//Delete Task
	r.POST("/delete-task/:id", func(c *gin.Context) {
		id := c.Param("id")
		user, err := c.Cookie("token")
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(user)
			oid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				log.Fatalf("Error converting to ObjectID: %v", err)
			}
			var task Task
			err = taskcollection.FindOneAndDelete(context.TODO(), bson.M{"_id": oid}).Decode(&task)
			c.Redirect(http.StatusFound, "/")
			if err != nil {
				log.Fatalf("Error finding task: %v", err)
			}
		}
	})
	//Finish Task
	r.POST("/finish-task/:id", func(c *gin.Context) {
		id := c.Param("id")
		user, err := c.Cookie("token")
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(user)
			oid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				log.Fatalf("Error converting to ObjectID: %v", err)
			}
			_, err = taskcollection.UpdateOne(context.TODO(), bson.M{"_id": oid}, bson.M{"$set": bson.M{"status": "finished"}})
			c.Redirect(http.StatusFound, "/")
			if err != nil {
				log.Fatalf("Error updating task status: %v", err)
			}
		}
	})
	//Unfinish Task
	r.POST("/unfinish-task/:id", func(c *gin.Context) {
		id := c.Param("id")
		user, err := c.Cookie("token")
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(user)
			oid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				log.Fatalf("Error converting to ObjectID: %v", err)
			}
			_, err = taskcollection.UpdateOne(context.TODO(), bson.M{"_id": oid}, bson.M{"$set": bson.M{"status": "not-finished"}})
			c.Redirect(http.StatusFound, "/")
			if err != nil {
				log.Fatalf("Error updating task status: %v", err)
			}
		}
	})
	//API
	r.GET("/api/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		var task Task
		err = taskcollection.FindOne(context.TODO(), bson.M{"userid": id}).Decode(&task)
		if err != nil {
			log.Fatalf("Error finding task: %v", err)
		}
		c.JSON(http.StatusOK, task)
	})
	r.Run(":5000")
}
