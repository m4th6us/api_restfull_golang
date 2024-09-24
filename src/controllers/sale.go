package controllers

import (
	"api_restfull/src/models"
	"encoding/json"
	"log"
	"net/http"
)

var Sales []models.Sale

func HandleReadSale(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	log.Println("Read Sales")

	json.NewEncoder(w).Encode(Sales)

}

func HandleCreateSale(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var sale models.Sale

	err := json.NewDecoder(r.Body).Decode(&sale)

	if err != nil{
		log.Fatal(err.Error())
	}

	Sales = append(Sales, sale)

	log.Println("Create Sale: ", sale)
	
	json.NewEncoder(w).Encode(sale)

}
