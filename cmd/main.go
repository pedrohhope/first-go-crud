package main

import (
	"first-go-crud/controllers"
	"first-go-crud/infra"
	"first-go-crud/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := infra.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/users", controllers.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUser(db)).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser(db)).Methods("DELETE")

	// Middleware
	router.Use(middlewares.JSONContentTypeMiddleware)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
