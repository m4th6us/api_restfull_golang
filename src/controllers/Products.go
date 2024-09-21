package controllers

import (
	"api_restfull/src/models"
	"encoding/json"
	"log"
	"net/http"
)

var Products []models.Products

func HandleReadProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	
	log.Println("Read Products")

	json.NewEncoder(w).Encode(Products)
	
}

func HandleCreateProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var product models.Products

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil{
		log.Fatal(err.Error())
	}

	Products = append(Products, product)
	
	log.Println("Create Products: ", product)

	json.NewEncoder(w).Encode(product)

}