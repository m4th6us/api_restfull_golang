package models

type Products struct {

	ProductId int `json:"product_id"`
	ProductName string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductPrice float64 `json:"product_price"`
	ProductStatus bool `json:"product_status"`
	
}

