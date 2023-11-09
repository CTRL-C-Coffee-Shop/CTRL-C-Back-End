package main

import (
	"fmt"
	"log"
	"net/http"

	"ctrl-c.com/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//authentication
	router.HandleFunc("/login", controllers.UserLogin).Methods("GET")
	router.HandleFunc("/logout", controllers.Logout).Methods("GET")

	//user

	//product
	router.HandleFunc("/product",controllers.Authenticate(controllers.GetProductsHandler, false)).Methods("GET")

	
	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080",router))
}