package controllers

import (
	"encoding/json"
	"net/http"

	"log"

	"api_restfull/src/models"
)

var Customers []models.Customer

func HandleReadCustomers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	
	log.Print("Read Customers")

	json.NewEncoder(w).Encode(Customers)

}

func HandleCreateCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var customer models.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		log.Fatal(err.Error())
	}

	Customers = append(Customers, customer)

	log.Print("Create Customer: ", customer)

	json.NewEncoder(w).Encode(customer)

}