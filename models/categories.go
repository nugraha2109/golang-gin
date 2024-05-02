package models

type Categories struct {
	Category_id   int64  `gorm:"primaryKey" json:"category_id"`
	Category_name string `gorm:"type:varchar(255)" json:"category_name"`
	Description   string `gorm:"type:varchar(255)" json:"description"`
}
