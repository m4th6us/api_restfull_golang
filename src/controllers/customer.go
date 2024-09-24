package controllers

import (
	"api_restfull/src/database"
	"api_restfull/src/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleFindOneCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

    log.Println("find one customer id ", params["customer_id"])

    var customer models.Customer
    result := database.DB.First(&customer, "id = ?", params["customer_id"])

    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(customer)
}


func HandleFindAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

    log.Println("find all customers")

    var customers []models.Customer
    result := database.DB.Find(&customers)
    
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(customers)
}


func HandleCreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

    var customer models.Customer
    err := json.NewDecoder(r.Body).Decode(&customer)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    log.Println("create new user: ", customer)

    result := database.DB.Create(&customer)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(customer)
}
