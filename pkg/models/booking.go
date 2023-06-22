package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	ID       int64
	BookedBy int64 `json:"booked_by"`
	User     User  `gorm:"foreignKey:BookedBy"`
}

func InitBooking(db *gorm.DB) {
	db.AutoMigrate(&Booking{})
}
