package models

type Product struct {
	Id          int    `gorm:"primaryKey;autoIncrement:true"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
}
