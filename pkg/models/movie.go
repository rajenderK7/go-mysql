package models

import "gorm.io/gorm"

type Movie struct {
	MovieID  int32  `json:"movie_id" gorm:"column:movie_id;primaryKey"`
	Name     string `json:"movie_name" gorm:"column:movie_name;not null"`
	Runtime  int    `json:"runtime" gorm:"column:runtime"`
	Director string `json:"director" gorm:"column:director;"`
}

func InitMovie(db *gorm.DB) {
	db.AutoMigrate(&Movie{})
}
