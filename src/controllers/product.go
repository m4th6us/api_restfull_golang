package controllers

import (
	"api_restfull/src/database"
	"api_restfull/src/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)


func HandleFindOneProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var product models.Product
	result := database.DB.Find(&product, "id = ?", params["product_id"])

	if result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusNotFound)
        return
	}

	json.NewEncoder(w).Encode(product)
}

func HandleFindAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    var products []models.Product
    result := database.DB.Find(&products)
    
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(products)
}

func HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    var product models.Product
    err := json.NewDecoder(r.Body).Decode(&product)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result := database.DB.Create(&product)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(product)
}