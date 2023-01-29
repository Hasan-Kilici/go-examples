package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://user:password@host:port"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("dbname").Collection("ogrenci")

	student := struct {
		Isim    string
		Soyisim string
		Sinif   string
		No      int
	}{
		Isim:    "Hasan",
		Soyisim: "KILICI",
		Sinif:   "12/D",
		No:      1200,
	}
	insertResult, err := collection.InsertOne(context.TODO(), student)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Eklenen veri: ", insertResult.InsertedID)
}
