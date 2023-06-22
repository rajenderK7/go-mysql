package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	ID     int64 `json:"booking_id" gorm:"column:booking_id;primaryKey"`
	UserID int64 `json:"booked_by"`
	User   User
}

func InitBooking(db *gorm.DB) {
	db.AutoMigrate(&Booking{})
}
