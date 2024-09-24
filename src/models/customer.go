package models

type Customer struct {

    Id        int    `gorm:"primaryKey;autoIncrement" json:"customer_id"`
    Name      string `json:"customer_name"`
    Age       int    `json:"customer_age"`
    Telephone string `json:"customer_telephone"`

}