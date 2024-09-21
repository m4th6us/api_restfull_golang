package models

type Customer struct {

	CustomerId int `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	CustomerAge int `json:"customer_age"`
	CustomerTelephone string `json:"customer_telephone"`

}