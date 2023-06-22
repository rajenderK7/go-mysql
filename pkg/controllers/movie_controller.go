package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rajenderK7/go-mysql/pkg/database"
	"github.com/rajenderK7/go-mysql/pkg/models"
)

func CreateMovie(ctx *fiber.Ctx) error {
	movie := models.Movie{}
	if err := ctx.BodyParser(&movie); err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"message": "Cannot parse movie data.",
			},
		)
		return err
	}
	db := database.GetDB()
	if err := db.Create(&movie).Error; err == nil {
		ctx.Status(http.StatusOK).JSON(
			&fiber.Map{
				"message": "Movie created succesfully.",
			},
		)
		return nil
	} else {
		ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "Something went wrong.",
			},
		)
		return err
	}
}

func GetMovies(ctx *fiber.Ctx) error {
	var movies []models.Movie
	db := database.GetDB()
	if err := db.Find(&movies).Error; err == nil {
		ctx.Status(http.StatusOK).JSON(
			&fiber.Map{
				"movies": movies,
			},
		)
		return nil
	} else {
		ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "Something went wrong.",
			},
		)
		return err
	}
}
