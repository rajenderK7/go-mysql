package models

import "gorm.io/gorm"

type Booking struct {
	BookingID int32 `json:"booking_id" gorm:"column:booking_id;primaryKey"`
	BookedBy  int32 `json:"booked_by" gorm:"column:user_id"`
	MovieID   int32 `json:"movie_id" gorm:"column:movie_id"`
	User      User  `gorm:"foreignKey:BookedBy;constraint:onUpdate:CASCADE,onDelete:SET NULL;"`
	Movie     Movie `gorm:"references:MovieID;constraint:onUpdate:CASCADE,onDelete:SET NULL;"`
}

func InitBooking(db *gorm.DB) {
	db.AutoMigrate(&Booking{})
}
