package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/y0gesh02/book_mana/database"
	"github.com/y0gesh02/book_mana/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



var collection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    var books []models.Book
    cur, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        log.Fatal(err)
    }
    defer cur.Close(context.Background())
    for cur.Next(context.Background()) {
        var item models.Book
        cur.Decode(&item)
        books = append(books, item)
    }
    if err := cur.Err(); err != nil {
        log.Fatal(err)
    }
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, err := primitive.ObjectIDFromHex(params["id"])
    if err != nil {
        log.Fatal(err)
    }
    var book models.Book
    err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&book)
    if err != nil {
        log.Fatal(err)
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    var book models.Book
    json.NewDecoder(r.Body).Decode(&book)
    result, err := collection.InsertOne(context.Background(), book)
    if err != nil {
        log.Fatal(err)
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(result)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}
    result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
    if err != nil {
        log.Fatal(err)
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(result)
}



func UpdateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}
    var book models.Book
    json.NewDecoder(r.Body).Decode(&book)
    result, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": book})
	if err != nil {
		log.Fatal(err)
	}
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(result)
}