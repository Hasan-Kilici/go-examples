package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserId string             `json:"userId"`
	Task   string             `json:"task"`
	Status string             `json:"status"`
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

	http.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
