package controllers

import (
	"api_restfull/src/models"
	"api_restfull/src/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var ArrayCustomers []models.Customer

func HandleFindOneCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	customer_id, _ := strconv.Atoi(params["customer_id"])

	log.Printf("Find customer id: %s", params["customer_id"])

	customer, err := utils.VerifyExistCustomerForFindOne(customer_id, ArrayCustomers)

	if err != nil{
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]*models.Customer{"ok": customer})

}

func HandleFindAllCustomers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	
	log.Print("Read Customers")

	if ArrayCustomers == nil{
		json.NewEncoder(w).Encode(map[string]string{"ok":"not exists customers"})
		return
	}

	json.NewEncoder(w).Encode(ArrayCustomers)

}

func HandleCreateCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var customer models.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		log.Fatal(err.Error())
	}

	customerExists := utils.VerifyExistCustomerForCreate(ArrayCustomers, customer)

	
	if customerExists{

		log.Print("Customer id already exists")

		w.WriteHeader(http.StatusConflict)

		message := "customer id: " + strconv.Itoa(customer.CustomerId) + " alread exists"

		json.NewEncoder(w).Encode(map[string]string{"error":message})

		return

	}

	ArrayCustomers = append(ArrayCustomers, customer)

	log.Print("Create Customer: ", customer)

	json.NewEncoder(w).Encode(customer)

}