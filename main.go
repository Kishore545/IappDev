package main

import (
	"log"
	"net/http"

	"IAppDev/database"
	"IAppDev/handlers"

	"github.com/gorilla/mux"
)

func main() {
	database.Init()

	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// Post routes
	r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
