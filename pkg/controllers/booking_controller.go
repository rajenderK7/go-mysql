package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rajenderK7/go-mysql/pkg/database"
	"github.com/rajenderK7/go-mysql/pkg/models"
)

func GetBookings(ctx *fiber.Ctx) error {
	var bookings []models.Booking
	db := database.GetDB()
	if err := db.Find(&bookings).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "Failed to fetch bookings",
			},
		)
	}
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"bookings": bookings,
		},
	)
}

func GetBookingsByUser(ctx *fiber.Ctx) error {
	id := ctx.Params("userID")
	userID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "Invalid user ID",
			},
		)
	}
	var bookings []models.Booking
	db := database.GetDB()
	if err := db.Select("booking_id, user_id, movie_id").Where("user_id = ?", int32(userID)).Preload("User").Preload("Movie").Find(&bookings).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"err":     err.Error,
				"message": "Failed to fetch bookings",
			},
		)
	}
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"bookings": bookings,
		},
	)
}
