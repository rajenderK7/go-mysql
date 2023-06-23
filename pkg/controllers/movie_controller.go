package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rajenderK7/go-mysql/pkg/database"
	"github.com/rajenderK7/go-mysql/pkg/models"
)

func CreateMovie(ctx *fiber.Ctx) error {
	movie := models.Movie{}
	if err := ctx.BodyParser(&movie); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "Cannot parse movie data",
			},
		)
	}
	db := database.GetDB()
	if err := db.Create(&movie).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "Failed to create movie",
			},
		)
	}
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Movie created successfully",
		},
	)
}

func GetMovies(ctx *fiber.Ctx) error {
	var movies []models.Movie
	db := database.GetDB()
	if err := db.Find(&movies).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "Could not fetch movies",
			},
		)
	}
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"movies": movies,
		},
	)
}

func GetMovieByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	movieID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "Invalid movie ID",
			},
		)
	}
	db := database.GetDB()
	movie := models.Movie{}
	if err := db.First(&movie, int32(movieID)).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "Could not find movie",
			},
		)
	}
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"movie": movie,
		},
	)
}
