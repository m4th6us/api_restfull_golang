package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"api_restfull/src/controllers"
	"api_restfull/src/database"
)

func main() {

	database.InitDB()

	r := mux.NewRouter()

	// custumers routes
	r.HandleFunc("/Customers", controllers.HandleFindAllCustomers).Methods("GET")
	r.HandleFunc("/Customers", controllers.HandleCreateCustomer).Methods("POST")
	r.HandleFunc("/Customers/{customer_id}", controllers.HandleFindOneCustomer).Methods("GET")

	// products routes
	r.HandleFunc("/Products", controllers.HandleFindAllProducts).Methods("GET")
	r.HandleFunc("/Products", controllers.HandleCreateProduct).Methods("POST")
	r.HandleFunc("/Products/{product_id}", controllers.HandleFindOneProduct).Methods("GET")


	// sales routes
	r.HandleFunc("/Sales", controllers.HandleReadSale).Methods("GET")
	r.HandleFunc("/Sales", controllers.HandleCreateSale).Methods("POST")

	http.ListenAndServe(":8090", r)
}