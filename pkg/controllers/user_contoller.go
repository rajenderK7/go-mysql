package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rajenderK7/go-mysql/pkg/database"
	"github.com/rajenderK7/go-mysql/pkg/models"
)

func CreateUser(ctx *fiber.Ctx) error {
	user := models.User{}
	if err := ctx.BodyParser(&user); err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"message": "Cannot parse user data.",
			},
		)
		return err
	}
	db := database.GetDB()
	if err := db.Create(&user).Error; err == nil {
		ctx.Status(http.StatusOK).JSON(
			&fiber.Map{
				"message": "User created succesfully.",
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

func GetUsers(ctx *fiber.Ctx) error {
	var users []models.User
	db := database.GetDB()
	if err := db.Find(&users).Error; err == nil {
		ctx.Status(http.StatusOK).JSON(
			&fiber.Map{
				"users": users,
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

func GetUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	user := &models.User{}
	db := database.GetDB()
	if err := db.Where("user_id = ?", userId).First(&user).Error; err == nil {
		ctx.Status(http.StatusOK).JSON(
			&fiber.Map{
				"user": user,
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
