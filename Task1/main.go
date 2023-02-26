package main

//Main file of our go project

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/y0gesh02/book_mana/routes"
)

func main(){
	
	r := mux.NewRouter()   //Initializing route
	routes.RegisterBookStoreRoutes(r)  //passing r to routes folder 
	http.Handle("/", r)  
	log.Fatal(http.ListenAndServe("localhost:8080", r)) //creating server
}