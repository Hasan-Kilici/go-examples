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

	taskcollection := client.Database("todolist").Collection("tasks")

	r := gin.Default()
	  r.GET("/api/tasks", func(c *gin.Context) {
	    var tasks []Task
	    cursor, err := taskcollection.Find(context.TODO(), bson.M{})
	    if err != nil {
	      log.Fatalf("Error finding tasks: %v", err)
	    }
	    for cursor.Next(context.TODO()) {
	      var task Task
	      if err := cursor.Decode(&task); err != nil {
		log.Fatalf("Error decoding task: %v", err)
	      }
	      tasks = append(tasks, task)
	    }
	    if err := cursor.Err(); err != nil {
	      log.Fatalf("Error iterating cursor: %v", err)
	    }
	    c.JSON(http.StatusOK, tasks)
	  })
  r.Run(":5000")
}



