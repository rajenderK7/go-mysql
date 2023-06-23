package models

import "gorm.io/gorm"

type User struct {
	UserID int32  `json:"user_id" gorm:"column:user_id;primaryKey"`
	Name   string `json:"name" gorm:"column:name;not null"`
	Phone  string `json:"phone" gorm:"column:phone;not null"`
}

func InitUser(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
