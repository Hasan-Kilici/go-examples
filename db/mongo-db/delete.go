package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("testdb").Collection("ogrenciler")

	deleteResult, err := collection.DeleteOne(context.TODO(),  bson.M{"no": 1200})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d kayıt silindi\n", deleteResult.DeletedCount)
}
