package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	ID       int64  `json:"movie_id"`
	Name     string `json:"name" gorm:"column:name;not null"`
	Runtime  int    `json:"runtime" gorm:"columne:runtime"`
	Director string `json:"director" gorm:"director"`
}

func InitMovie(db *gorm.DB) {
	db.AutoMigrate(&Movie{})
}
