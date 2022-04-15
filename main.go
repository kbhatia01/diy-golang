package main

import (
	connection "diy/database"
	"diy/pkg/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("hey")
	db := connection.Init()
	h := handlers.NewHandler(db)
	router := mux.NewRouter()
	router.HandleFunc("/products", h.GetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", h.AddProduct).Methods(http.MethodPost)
	router.HandleFunc("/products/{id}", h.GetProductById).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)
}
