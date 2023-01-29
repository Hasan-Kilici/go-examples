package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	ogrenciCollection := client.Database("veritabaniAdi").Collection("ogrenci")

	updateResult, err := ogrenciCollection.UpdateOne(
		context.TODO(),
		bson.M{"no": 1200},
		bson.D{
			{"$set", bson.D{
				{"isim", "Hasan"},
				{"soyisim", "KILICI"},
				{"sinif", "12/D"},
			}},
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d satır düzenlendi\n", updateResult.ModifiedCount)
}
