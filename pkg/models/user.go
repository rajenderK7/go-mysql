package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    int64  `json:"user_id"`
	Name  string `json:"name" gorm:"column:name;not null"`
	Phone string `json:"phone" gorm:"column:phone;not null"`
}

func InitUser(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
