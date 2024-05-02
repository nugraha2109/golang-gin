package models

type Barangs struct {
	Id          int64        `gorm:"primaryKey" json:"id"`
	ProductName string       `gorm:"type:varchar(255)" json:"product_name"`
	Category_id uint         `gorm:"type:id(20)" json:"category_id"`
	Unit        string       `gorm:"type:varchar(255)" json:"unit"`
	Price       string       `gorm:"type:varchar(255)" json:"price"`
	Barangs     []Categories `gorm:"Foreignkey:category_id;association_foreignkey:id;" json:"categories"`
}
