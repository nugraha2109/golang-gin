package models

type Prods struct {
	Product_id   int64  `gorm:"primaryKey" json:"product_id"`
	Product_name string `gorm:"type:varchar(255)" json:"product_name"`
	Supplier_id  int64  `gorm:"type:integer" json:"supplier_id"`
	Category_id  int64  `gorm:"type:integer" json:"category_id"`
	Unit         string `gorm:"type:varchar(255)" json:"unit"`
	Price        string `gorm:"type:varchar(255)" json:"price"`
}
