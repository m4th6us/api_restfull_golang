package models

type Sales struct {

	SaleId int `json:"sale_id"`
	SaleValue float64 `json:"sale_value"`
	SaleStatus bool `json:"sale_status"`
	SaleDate string `json:"sale_date"`
	Customer Customer
	Products []Products

}