package models

type Product struct {

    Id           int     `gorm:"primaryKey;autoIncrement" json:"product_id"`
    Name         string  `json:"product_name"`
    Description  string  `json:"product_description"`
    Price        float64 `json:"product_price"`
    Status       bool    `json:"product_status"`

}
