package models

import "time"

type Sale struct {

    Id     int       `json:"sale_id" gorm:"primaryKey;autoIncrement"`
    Value  float64   `json:"sale_value"`
    Status bool      `json:"sale_status"`
    Date   time.Time `json:"sale_date"`
    
    CustomerId int       `json:"customer_id"`
    Customer   Customer  `gorm:"foreignKey:CustomerId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"customer"`
    
    Products    []Product `gorm:"many2many:sale_product;" json:"products"`

}



