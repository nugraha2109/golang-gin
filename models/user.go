package models

type Users struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"type:varchar(255))" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	APIkey   string `gorm:"type:varchar(255)" json:"apikey"`
}
