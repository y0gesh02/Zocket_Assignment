package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//MongoDb := "mongodb+srv://yogesh02:Devil1234@cluster0.euqff.mongodb.net/?retryWrites=true&w=majority"

	//MongoDb:=os.Getenv("MONGODB_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://yogesh02:1234@cluster0.xu1dpbn.mongodb.net/?retryWrites=true&w=majority")) //connecting
	if err != nil {
		log.Fatal(err)
	}
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //timeout
	defer cancel()
	err = client.Connect(ctx) //connecting to db
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}
var Client *mongo.Client = DBinstance() //client instance

// opening/access  collection
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}