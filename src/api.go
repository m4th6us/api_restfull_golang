package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"api_restfull/src/controllers"
)

func main() {
	r := mux.NewRouter()

	// custumers routes
	r.HandleFunc("/Customers", controllers.HandleReadCustomers).Methods("GET")
	r.HandleFunc("/Customers", controllers.HandleCreateCustomer).Methods("POST")

	// products routes
	r.HandleFunc("/Products", controllers.HandleReadProducts).Methods("GET")
	r.HandleFunc("/Products", controllers.HandleCreateProducts).Methods("POST")

	// sales routes
	r.HandleFunc("/Sales", controllers.HandleReadSale).Methods("GET")
	r.HandleFunc("/Sales", controllers.HandleCreateSale).Methods("POST")

	http.ListenAndServe(":8090", r)
}